package main

import (
	"log"
	"time"
)

func main() {

	for true {
		ip, err := GetPublicIp()
		if err != nil {
			log.Printf("Could not fetch public IP because of an error: %s", err)
			continue
		}
		log.Printf("Got public IP %s", ip)

		doesExist, record := DoesRecordAlreadyExist()

		if doesExist {
			UpdateRecord(ip, record)
		} else {
			CreateNewRecord(ip)
		}

		log.Printf("Waiting for next update...")
		time.Sleep(1200 * time.Second)
	}
}
