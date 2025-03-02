package main

import (
	"errors"

	"gopkg.in/ini.v1"
)

type Config struct {
	Interface string `ini:"interface"`
	Clients   map[string]ClientConfig
}

type ClientConfig struct {
	Hostname string `ini:"hostname"`
	MacAddr  string `ini:"macaddress"`
}

func ParseConfigFile(filename string) (*Config, error) {
	data, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}

	cfg := &Config{Clients: make(map[string]ClientConfig)}

	err = data.Section("").MapTo(cfg)
	if err != nil {
		return nil, err
	}
	if cfg.Interface == "" {
		return nil, errors.New("Config does not contain an interface")
	}

	for _, section := range data.Sections() {
		if section.Name() == ini.DefaultSection {
			continue
		}

		var client ClientConfig
		err := section.MapTo(&client)
		if err != nil {
			return nil, err
		}

		if client.Hostname == "" {
			return nil, errors.New(section.Name() + " does not contain a hostname")
		}
		if client.MacAddr == "" {
			return nil, errors.New(section.Name() + " does not contain a mac address")
		}

		cfg.Clients[section.Name()] = client
	}

	return cfg, nil
}
