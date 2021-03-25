package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	ds, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(ds), "interface(s):")

	for _, d := range ds {
		for _, a := range d.Addresses {
			fmt.Printf("interface = %s ; IP = %s ; mask = %s\n", d.Name, a.IP, a.Netmask)
		}
	}
}
