package pkg

import "time"

type UpdateDnsResponse struct {
	Success  bool     `json:"success"`
	Errors   []Error  `json:"errors"`
	Messages []string `json:"messages"`
	Result   DnsEntry `json:"result"`
}

type ListZonesResponse struct {
	Success    bool      `json:"success"`
	Errors     []Error   `json:"errors"`
	Messages   []Message `json:"messages"`
	Result     []Zone    `json:"result"`
	ResultInfo Info      `json:"result_info"`
}

type ListDnsResponse struct {
	Success    bool       `json:"success"`
	Errors     []Error    `json:"errors"`
	Messages   []Message  `json:"messages"`
	Result     []DnsEntry `json:"result"`
	ResultInfo Info       `json:"result_info"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DnsRecord struct {
	Content string   `json:"content"`
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Proxied bool     `json:"proxied"`
	Comment string   `json:"comment"`
	Tags    []string `json:"tags"`
	Ttl     int      `json:"ttl"`
}

type Zone struct {
	Account             any       `json:"account"`
	ActivatedOn         time.Time `json:"activated_on"`
	CreatedOn           time.Time `json:"created_on"`
	DevelopmentMode     int       `json:"development_mode"`
	ID                  string    `json:"id"`
	Meta                Meta      `json:"meta"`
	ModifiedOn          time.Time `json:"modified_on"`
	Name                string    `json:"name"`
	OriginalDNSHost     string    `json:"original_dnshost"`
	OriginalNameServers []string  `json:"original_name_servers"`
	OriginalRegistrar   string    `json:"original_registrar"`
	Owner               any       `json:"owner"`
	VanityNameServers   []string  `json:"vanity_name_servers"`
}

type DnsEntry struct {
	Content    string    `json:"content"`
	Name       string    `json:"name"`
	Proxied    bool      `json:"proxied"`
	Type       string    `json:"type"`
	Comment    string    `json:"comment"`
	CreatedOn  time.Time `json:"created_on"`
	ID         string    `json:"id"`
	Locked     bool      `json:"locked"`
	Meta       Meta      `json:"meta"`
	ModifiedOn time.Time `json:"modified_on"`
	Proxiable  bool      `json:"proxiable"`
	Tags       []string  `json:"tags"`
	TTL        int       `json:"ttl"`
	ZoneID     string    `json:"zone_id"`
	ZoneName   string    `json:"zone_name"`
}

type Meta struct {
	AutoAdded bool   `json:"auto_added"`
	Source    string `json:"source"`
}

type Info struct {
	Count      int `json:"count"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
}
