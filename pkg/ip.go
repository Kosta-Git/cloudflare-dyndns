package pkg

import (
	"io"
	"log"
	"net/http"
	"net/netip"
	"strings"
)

const ipServiceUrl = "http://checkip.amazonaws.com/"

func FetchPublicIp() netip.Addr {
	res, err := http.Get(ipServiceUrl)
	if err != nil {
		log.Fatalf("Error fetching public IP: %v\n", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return netip.MustParseAddr(strings.Trim(string(body), "\n"))
}
