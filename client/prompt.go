package main

import (
	"fmt"
	"log"

	tui "github.com/manifoldco/promptui"
)

func (rt *Runtime) PromptMainMenu() {
	var items []string
	for clientID := range rt.Clients {
		items = append(items, clientID)
	}
	items = append(items, "EXIT")
	prompt := tui.Select{
		Label: "Which client would you like to manage",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	if result == "EXIT" {
		return
	}

	rt.PromptManageClient(result)
	rt.PromptMainMenu()
}

func (rt *Runtime) PromptManageClient(clientID string) {
	c := rt.Clients[clientID]
	fmt.Println(
		"Current State:",
		ClientStateToString[c.Runtime.State],
	)
	fmt.Println(
		"Last State:",
		ClientStateToString[c.Runtime.LastState],
	)
	fmt.Println()

	if c.Runtime.LeaseTime > 0 {
		secsLeft := c.Runtime.LeaseTime - c.Runtime.LeaseSecsElapsed
		mins := secsLeft / 60
		secs := secsLeft % 60
		fmt.Printf(
			"Lease expires in: %02d:%02d\n",
			mins, secs,
		)
		fmt.Println()
	}

	fmt.Println(
		"Hostname:",
		c.Hostname,
	)
	fmt.Println(
		"Mac Address:",
		c.MacAddr.String(),
	)
	fmt.Println()

	if ClientStateShouldContainNetwork(c.Runtime.State) {
		fmt.Println(
			"Network:",
			c.Runtime.Network.String(),
		)
		fmt.Println(
			"IP Address:",
			c.Runtime.IP.String(),
		)
		fmt.Println(
			"Default Gateway:",
			c.Runtime.DefaultGateway.String(),
		)
		fmt.Println(
			"DHCP Server IP:",
			c.Runtime.DHCPServerIP.String(),
		)
		fmt.Println()
	}

	mgmtFuncs, labels := rt.GetManagementOptions()
	labels = append(labels, "EXIT")

	prompt := tui.Select{
		Label: "What would you like to do",
		Items: labels,
	}
	index, result, _ := prompt.Run()
	if result == "EXIT" {
		return
	}

	mgmtFuncs[index](rt, clientID)

	rt.PromptManageClient(clientID)
}
