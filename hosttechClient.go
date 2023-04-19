package main

import (
	"context"
	"github.com/libdns/hosttech"
	"github.com/libdns/libdns"
	"log"
	"os"
)

func UpdateHosttechRecords(ip string) {
	zone := os.Getenv("ZONE")
	domain := os.Getenv("DOMAIN")

	log.Printf("Updating A record for zone %s, domain %s with IP %s", zone, domain, ip)

	provider := hosttech.Provider{
		APIToken: os.Getenv("API_KEY"),
	}

	_, err := provider.AppendRecords(context.Background(), zone, []libdns.Record{
		{
			ID:       "",
			Type:     "A",
			Name:     domain,
			Value:    ip,
			TTL:      3600,
			Priority: 0,
		},
	})
	if err != nil {
		log.Printf("Could not update IP because of an error %s", err)
	}
}
