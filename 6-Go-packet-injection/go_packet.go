package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	// find all Device
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	for _, device := range devices {
		fmt.Println("Name:", device.Name)
		fmt.Println("Description", device.Description)
		fmt.Println("Devices address", device.Description)

		for _, address := range device.Addresses {
			fmt.Println("- IP address", address.IP)
			fmt.Println("- Subnet mask:", address.Netmask)
		}
	}

}
