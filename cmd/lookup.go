package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/extractip/ipwho/internal"
	"io"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Lookup exit IP address",
	Long:  "Query ExtractIP API to resolve the exit IP address.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return lookupHost()
	},
}

func init() {
	lookupCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed information in table format")
	lookupCmd.PersistentFlags().BoolVar(&jsonOut, "json", false, "Output the result in JSON format")

	rootCmd.AddCommand(lookupCmd)
}

func lookupHost() error {
	url := "https://api.extractip.com/geolocate"
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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: %s", string(body))
	}

	var parsed internal.LookupResponse
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

	// Default: just IP address
	if parsed.ExitIP != "" {
		fmt.Println(parsed.ExitIP)
	} else {
		fmt.Println("IP address not found in response")
	}
	return nil
}
