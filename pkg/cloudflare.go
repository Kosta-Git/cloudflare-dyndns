package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const cloudflareApiUrl = "https://api.cloudflare.com/client/v4/"

func ListZones(apiKey string) []Zone {
	req, err := http.NewRequest("GET", cloudflareApiUrl+"zones", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error: %v\n", res.Status)
	}

	var response ListZonesResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	if !response.Success {
		log.Fatalf("Error: %v\n", response.Errors)
	}

	if response.ResultInfo.TotalCount > response.ResultInfo.Count {
		log.Print("Zone pagination not implemented yet")
	}

	return response.Result
}

func FindZoneId(zone, apiKey string) (string, error) {
	zones := ListZones(apiKey)
	for _, z := range zones {
		if z.Name == zone {
			return z.ID, nil
		}
	}
	return "", errors.New("Zone not found")
}

func ListRecords(zoneId string, apiKey string) []DnsEntry {
	req, err := http.NewRequest("GET", cloudflareApiUrl+"zones/"+zoneId+"/dns_records", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error: %v\n", res.Status)
	}

	var response ListDnsResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	if !response.Success {
		log.Fatalf("Error: %v\n", response.Errors)
	}

	return response.Result
}

func UpdateDnsEntry(zoneId, entryId string, entry DnsRecord, apiKey string) bool {
	jsonData, err := json.Marshal(entry)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v\n", err)
	}

	req, err := http.NewRequest("PUT", cloudflareApiUrl+"zones/"+zoneId+"/dns_records/"+entryId, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error: %v\n", res.Status)
	}

	var response UpdateDnsResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	if !response.Success {
		log.Fatalf("Error: %v\n", response.Errors)
	}

	return true
}
