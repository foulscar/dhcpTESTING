#!/bin/bash
blue=$(tput setaf 4)
magenta=$(tput setaf 5)
red=$(tput setaf 1)
normal=$(tput sgr0)

printf "${red}It is recommended to run this script in a sandbox environment"
printf " as it uses podman in the host network mode which could cause issues on the host machine\n\n"

printf "${blue}This script requires sudo\n"
printf "Please enter your password\n\n${normal}"

if ! sudo true; then
        exit 1
fi

if ! command -v podman &> /dev/null; then
        printf "${red}podman is not installed. Exiting\n${normal}"
        exit 1
fi

printf "${blue}Removing bridge/veth interfaces"
printf "${blue}' and creating again\n${normal}"

printf "${red}Removing...\n${normal}"
sudo ip link del dhcp-test-br
sudo ip link del dhcp-test-veth0
sudo ip link del dhcp-test-veth1
printf "${blue}Creating...\n\n${normal}"
sudo ip link add dhcp-test-br type bridge
sudo ip link add dhcp-test-veth0 type veth peer name dhcp-test-veth1
sudo ip link set dhcp-test-veth0 master dhcp-test-br
sudo ip addr add 10.80.54.1/24 dev dhcp-test-br
sudo ip link set dev dhcp-test-br up
sudo ip link set dev dhcp-test-veth0 up
sudo ip link set dev dhcp-test-veth1 up

printf "${blue}Building podman image\n${normal}"
sudo podman image build -t dhcp-test:latest .

printf "${blue}A container with the name '${magenta}dhcp-test${blue}' is about to start\n"
printf "You will be able to access it via the '${magenta}dhcp-test-veth1${blue}' interface\n"
printf "isc-dhcp-client is recommended\n"
printf "Ctrl+C will kill it and clean up this environment for you\n\n"
printf "Press any key to start the container\n\n${normal}"

read -n 1 -s

printf "${blue}Running podman container\n${normal}"
sudo podman run \
        -it \
        --privileged \
        --replace \
        --network host \
        --name dhcp-test \
        dhcp-test:latest

printf "\n\n${blue}You may need to enter your password again\n"
printf "${blue}Failure to do so will exit this script and not clean up\n\n${normal}"

if ! sudo true; then
        exit 1
fi

printf "${red}Removing container\n${normal}"
sudo podman rm dhcp-test
printf "${red}Deleting bridge/veth interfaces\n${normal}"
sudo ip link del dhcp-test-br
sudo ip link del dhcp-test-veth0
