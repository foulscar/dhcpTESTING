#!/bin/bash

# Forward traffic and act as a router
iptables -A FORWARD -i dhcp-test-br -o dhcp-test-br -j ACCEPT

# Allow incoming DHCP requests (clients -> DHCP server)
iptables -A INPUT -i dhcp-test-br -p udp --dport 67 -j ACCEPT
iptables -A INPUT -i dhcp-test-br -p udp --dport 68 -j ACCEPT

# Allow DHCP replies (server -> clients)
iptables -A OUTPUT -o dhcp-test-br -p udp --sport 67 -j ACCEPT
iptables -A OUTPUT -o dhcp-test-br -p udp --sport 68 -j ACCEPT

# Allow ICMP (Ping)
iptables -A INPUT -i dhcp-test-br -p icmp -j ACCEPT
iptables -A OUTPUT -o dhcp-test-br -p icmp -j ACCEPT

# Drop all other traffic
iptables -A INPUT -i dhcp-test-br -j DROP
iptables -A OUTPUT -o dhcp-test-br -j DROP

dhcpd -f -d -cf /etc/dhcp/dhcpd.conf dhcp-test-br
