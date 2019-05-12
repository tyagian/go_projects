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
	for i in devces {
		fmt.Println(i)
	}

}