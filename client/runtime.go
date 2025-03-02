package main

import (
	"fmt"

	"github.com/foulscar/dhcp"
)

type Runtime struct {
	Clients map[string]Client
	Conn    *dhcp.DHCPConn
}

func (c Config) BuildRuntime() (*Runtime, error) {
	rt := &Runtime{
		Clients: make(map[string]Client),
	}

	conn, err := dhcp.NewDHCPConn(c.Interface, 67, 68)
	if err != nil {
		return nil, fmt.Errorf("error creating connection: %w", err)
	}
	rt.Conn = conn

	for id, clientConfig := range c.Clients {
		client, err := clientConfig.BuildClient()
		if err != nil {
			return nil, fmt.Errorf("error building client '%s' runtime: %w", id, err)
		}

		rt.Clients[id] = *client
	}

	return rt, nil
}
