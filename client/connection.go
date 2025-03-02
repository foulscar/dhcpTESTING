package main

import (
	"fmt"

	"github.com/foulscar/dhcp"
)

func (rt *Runtime) Listen() {
	buffer := make([]byte, 1500)
	for {
		msg, err := rt.ListenForMessage(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(*msg)
	}
}

func (rt *Runtime) ListenForMessage(buffer []byte) (*dhcp.Message, error) {
	n, err := rt.Conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	if !dhcp.IsEncodedMessage(buffer[:n]) {
		return nil, nil
	}

	msg, err := dhcp.MarshalMessage(buffer[:n])
	if err != nil {
		return nil, err
	}

	return msg, nil
}
