package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	iface, err := FindMyInterface()
	if err != nil {
		panic(err)
	}

	handle, err := pcap.OpenLive(iface, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	fmt.Println("Sniff traffic...")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		dnsLayer := packet.Layer(layers.LayerTypeDNS)
		if dnsLayer == nil {
			continue
		}

		dns, _ := dnsLayer.(*layers.DNS)
		for _, q := range dns.Questions {
			src := packet.NetworkLayer().NetworkFlow().Src().String()
			fmt.Printf("🔍 [%s] DNS query: %s\n", src, string(q.Name))
		}
	}
}
