package main

import (
	"log"
	"time"
	"strings"
	"os"
)

func main() {

	for true {
		ip, err := GetPublicIp()
		if err != nil {
			log.Printf("Could not fetch public IP because of an error: %s", err)
			continue
		}
		log.Printf("Got public IP %s", ip)

		domains := strings.Split(os.Getenv("DOMAINS"), (","));
		for _, domain := range domains {
			doesExist, record := DoesRecordAlreadyExist(domain)

			if doesExist {
				UpdateRecord(ip, record)
			} else {
				CreateNewRecord(ip, domain)
			}
		}
		
		log.Printf("Waiting for next update...")
		time.Sleep(1200 * time.Second)
	}
}
