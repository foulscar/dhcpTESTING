package main

import (
	"fmt"

	"github.com/foulscar/dhcp"
)

type ManagementFunc func(*Runtime, string) error

func (rt *Runtime) GetManagementOptions() ([]ManagementFunc, []string) {
	funcs := []ManagementFunc{ManagementTestFunc}
	labels := []string{"Test"}

	return funcs, labels
}

func ManagementTestFunc(rt *Runtime, clientID string) error {
	msg := dhcp.NewEmptyMessage()
	msg.Flags = 0x8000
	msg.BOOTPMessageType = dhcp.BOOTPMessageTypeRequest
	msg.HardwareAddrType = dhcp.HardwareAddrTypeEthernet
	msg.HardwareAddrLen = uint8(6)
	msg.TransactionID = 0xFFFFFFFF
	msg.ClientHardwareAddr = rt.Clients[clientID].MacAddr
	opt, err := dhcp.NewOptionMessageType(dhcp.OptionMessageTypeCodeDISCOVER)
	msg.Options.Add(opt)

	data := msg.Unmarshal()
	_, err = rt.Conn.Write(data)
	if err != nil {
		return err
	}

	fmt.Println("Test:", len(data))

	return nil
}
