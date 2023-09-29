package pkg

import (
	"errors"
	"log"
	"net/netip"
)

var (
	ZoneId       string     = ""
	LastIp       netip.Addr = netip.IPv4Unspecified()
	LastDnsEntry DnsEntry   = DnsEntry{}
)

func UpdateDynDns(zone, subZone, apiKey string) {
	zoneId, err := getZoneId(zone, apiKey)
	if err != nil {
		log.Fatalf("Error getting zone id: %v\n", err)
	}
	publicIp := FetchPublicIp()
	if LastIp == publicIp {
		log.Print("IP address has not changed")
		return
	}
	if LastDnsEntry.ID == "" {
		LastDnsEntry, err = getDnsEntry(zoneId, subZone, apiKey)
		if err != nil {
			log.Fatalf("Error getting DNS entry: %v\n", err)
		}
	}
	if LastDnsEntry.Content == publicIp.String() {
		log.Print("DNS entry is up to date")
		return
	}
	LastDnsEntry, err = getDnsEntry(zoneId, subZone, apiKey)
	if err != nil {
		createEntry(zoneId, subZone, publicIp.String(), apiKey)
		return
	}
	if LastDnsEntry.Content != publicIp.String() {
		updateEntry(zoneId, publicIp.String(), apiKey, LastDnsEntry)
		return
	}
}

func createEntry(zoneId, subZone, ip, apiKey string) {
	log.Printf("DNS Entry does not exist, creating it with ip: %s\n", FetchPublicIp().String())
	CreateDnsEntry(zoneId, DnsRecord{
		Content: ip,
		Name:    subZone,
		Type:    "A",
		Proxied: false,
		Comment: "DynDNS by cf-dyndns",
		Tags:    nil,
		Ttl:     1,
	}, apiKey)
}

func updateEntry(zoneId, ip, apiKey string, dnsEntry DnsEntry) {
	log.Printf("IP address has changed, updating DNS entry [current: %s, new: %s]\n", dnsEntry.Content, ip)
	UpdateDnsEntry(zoneId, dnsEntry.ID, DnsRecord{
		Content: ip,
		Name:    dnsEntry.Name,
		Type:    dnsEntry.Type,
		Proxied: dnsEntry.Proxied,
		Comment: "DynDNS by cf-dyndns",
		Tags:    nil,
		Ttl:     dnsEntry.TTL,
	}, apiKey)
}

func getDnsEntry(zoneId, dnsEntryName, apiKey string) (DnsEntry, error) {
	dnsEntries := ListRecords(dnsEntryName, "A", zoneId, apiKey)
	for _, dnsEntry := range dnsEntries {
		if dnsEntry.Name == dnsEntryName && dnsEntry.Type == "A" {
			return dnsEntry, nil
		}
	}
	return DnsEntry{}, errors.New("DNS entry not found")
}

func getZoneId(zone, apiKey string) (string, error) {
	if ZoneId == "" {
		var err error
		ZoneId, err = FindZoneId(zone, apiKey)
		if err != nil {
			return "", err
		}
	}
	return ZoneId, nil
}
