package main

import (
	"fmt"
	"log"
	"net"
)

func FindMyInterface() (string, error) {
	i, err := net.Interfaces()
	if err != nil {
		log.Fatalf("We have error: %w ", err)
	}
	for _, f := range i {
		if f.Flags&net.FlagUp != 0 && f.Flags&net.FlagLoopback == 0 {
			addrs, _ := f.Addrs()
			for _, a := range addrs {
				ipnet, ok := a.(*net.IPNet)
				if !ok || ipnet.IP.IsLoopback() {
					continue
				}
				ip := ipnet.IP.To4()
				if ip == nil {
					continue
				}
				return f.Name, nil
			}
		}
	}
	return "", fmt.Errorf("No interfaces")
}
