package internal

type LookupResponse struct {
	ExitIP        string   `json:"exit_ip"`
	Capital       string   `json:"capital"`
	CountryCode   string   `json:"country_code"`
	CountryName   string   `json:"country_name"`
	CountryRegion string   `json:"country_region"`
	TimeZone      string   `json:"time_zone"`
	DomainName    string   `json:"domain_name"`
	Currency      string   `json:"currency"`
	FlagURL       string   `json:"flag_url"`
	Domains       []string `json:"domains"`
	CallCodes     []string `json:"call_codes"`
}
