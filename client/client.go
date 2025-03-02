package main

import (
	"errors"
	"net"
)

type Client struct {
	Hostname string
	MacAddr  net.HardwareAddr
	Runtime  ClientRuntime
}

type ClientRuntime struct {
	State            ClientState
	LastState        ClientState
	IP               net.IP
	Network          *net.IPNet
	DefaultGateway   net.IP
	DHCPServerIP     net.IP
	LeaseTime        int
	LeaseSecsElapsed int
}

func (cc ClientConfig) BuildClient() (*Client, error) {
	client := &Client{
		Runtime: ClientRuntime{
			State:     ClientStateINIT,
			LastState: ClientStateNIL,
			LeaseTime: 255,
		},
	}

	client.Hostname = cc.Hostname

	var err error
	client.MacAddr, err = net.ParseMAC(cc.MacAddr)
	if err != nil {
		return nil, errors.New("Failed to parse macaddress. " + err.Error())
	}

	return client, nil
}
