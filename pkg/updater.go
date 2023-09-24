package pkg

import (
	"errors"
	"log"
)

func UpdateDynDns(zone, subZone, apiKey string) {
	zoneId, err := FindZoneId(zone, apiKey)
	if err != nil {
		log.Fatalf("Error finding zone ID: %v\n", err)
	}
	dnsEntry, err := getDnsEntry(zoneId, subZone, apiKey)
	if err != nil {
		log.Fatalf("Error finding DNS entry: %v\n", err)
	}
	publicIp := FetchPublicIp()

	if dnsEntry.Content == publicIp.String() {
		log.Print("IP address has not changed")
	} else {
		log.Printf("IP address has changed, updating DNS entry [current: %s, new: %s]\n", dnsEntry.Content, publicIp.String())
		UpdateDnsEntry(zoneId, dnsEntry.ID, DnsRecord{
			Content: publicIp.String(),
			Name:    dnsEntry.Name,
			Type:    dnsEntry.Type,
			Proxied: dnsEntry.Proxied,
			Comment: "DynDNS by cf-dyndns",
			Tags:    nil,
			Ttl:     dnsEntry.TTL,
		}, apiKey)
	}
}

func getDnsEntry(zoneId, dnsEntryName, apiKey string) (DnsEntry, error) {
	dnsEntries := ListRecords(zoneId, apiKey)
	for _, dnsEntry := range dnsEntries {
		if dnsEntry.Name == dnsEntryName && dnsEntry.Type == "A" {
			return dnsEntry, nil
		}
	}
	return DnsEntry{}, errors.New("DNS entry not found")
}
