package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

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

var verbose bool
var jsonOut bool

var rootCmd = &cobra.Command{
	Use:   "ipwho [ip]",
	Short: "ipwho - CLI tool to get the country of an IP address",
	Long:  "ipwho is a lightweight CLI tool, similar to whois, but focused on detecting the country (and extended details) of an IP address via ExtractIP API.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ip := args[0]
		return lookupIP(ip)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed information in table format")
	rootCmd.PersistentFlags().BoolVar(&jsonOut, "json", false, "Output the result in JSON format")
}

func lookupIP(ip string) error {
	url := "https://api.extractip.com/geolocate/" + ip
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("API error: %s", string(body))
	}

	var parsed LookupResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return err
	}

	// JSON output
	if jsonOut {
		pretty, _ := json.MarshalIndent(parsed, "", "  ")
		fmt.Println(string(pretty))
		return nil
	}

	// Verbose table output
	if verbose {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		fmt.Fprintf(w, "Exit IP:\t%s\n", parsed.ExitIP)
		fmt.Fprintf(w, "Country:\t%s (%s)\n", parsed.CountryName, parsed.CountryCode)
		fmt.Fprintf(w, "Region:\t%s\n", parsed.CountryRegion)
		fmt.Fprintf(w, "Capital:\t%s\n", parsed.Capital)
		fmt.Fprintf(w, "Time Zone:\t%s\n", parsed.TimeZone)
		fmt.Fprintf(w, "Currency:\t%s\n", parsed.Currency)
		fmt.Fprintf(w, "Domain:\t%s\n", parsed.DomainName)
		fmt.Fprintf(w, "Flag URL:\t%s\n", parsed.FlagURL)

		if len(parsed.Domains) > 0 {
			fmt.Fprintf(w, "Domains:\t%s\n", strings.Join(parsed.Domains, ", "))
		}
		if len(parsed.CallCodes) > 0 {
			fmt.Fprintf(w, "Call Codes:\t%s\n", strings.Join(parsed.CallCodes, ", "))
		}

		w.Flush()
		return nil
	}

	// Default: just country name
	if parsed.CountryName != "" {
		fmt.Println(parsed.CountryName)
	} else {
		fmt.Println("Country not found in response")
	}
	return nil
}
