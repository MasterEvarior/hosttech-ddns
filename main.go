package main

import (
	"log"
	"time"
)

func main() {

	for true {
		time.Sleep(1200 * time.Second)
		ip, err := GetPublicIp()
		if err != nil {
			log.Printf("Could not fetch public IP because of an error: %s", err)
			continue
		}

		log.Printf("Got public IP %s", ip)

		UpdateHosttechRecords(ip)
	}
}
