package main

import (
	"context"
	"github.com/libdns/hosttech"
	"github.com/libdns/libdns"
	"log"
	"os"
)

func DoesRecordAlreadyExist() (bool, libdns.Record) {
	record, err := getRecord()

	if err == nil && record.ID != "" {
		return true, record
	} else {
		return false, libdns.Record{}
	}
}

func UpdateRecord(ip string, existingRecord libdns.Record) {
	log.Printf("Updating A record for zone %s, domain %s with IP %s", os.Getenv("ZONE"), os.Getenv("DOMAIN"), ip)
	provider := hosttech.Provider{APIToken: os.Getenv("API_KEY")}

	_, err := provider.SetRecords(context.Background(), os.Getenv("ZONE"), []libdns.Record{
		{
			ID:       existingRecord.ID,
			Type:     existingRecord.Type,
			Name:     existingRecord.Name,
			Value:    ip,
			TTL:      3600,
			Priority: 0,
		},
	})

	if err != nil {
		log.Printf("Could not update IP because of an error %s", err)
	}
}

func CreateNewRecord(ip string) {
	log.Printf("Creating a new A record for zone %s, domain %s with IP %s", os.Getenv("ZONE"), os.Getenv("DOMAIN"), ip)

	provider := hosttech.Provider{APIToken: os.Getenv("API_KEY")}
	_, err := provider.AppendRecords(context.Background(), os.Getenv("ZONE"), []libdns.Record{
		{
			ID:       "",
			Type:     "A",
			Name:     os.Getenv("DOMAIN"),
			Value:    ip,
			TTL:      3600,
			Priority: 0,
		},
	})
	if err != nil {
		log.Printf("Could not create new record because of an error %s", err)
	}
}

func getRecord() (libdns.Record, error) {
	provider := hosttech.Provider{APIToken: os.Getenv("API_KEY")}
	records, err := provider.GetRecords(context.Background(), os.Getenv("ZONE"))

	if err != nil {
		log.Printf("Could not fetch existing records from API because of an error %s", err)
		return libdns.Record{}, err
	}

	for _, record := range records {
		if record.Type == "A" && record.Name == os.Getenv("DOMAIN") {
			return record, nil
		}
	}

	return libdns.Record{}, nil
}
