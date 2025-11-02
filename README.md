# ipwho

`ipwho` is a lightweight CLI tool for Unix-like systems, similar to `whois`, but focused on retrieving the **country** and extended details of an IP address.

## 🚀 Features

* Quickly detect the country of an IP address
* Verbose mode (`-v`) to show extended details:

    * Country, country code, region
    * Capital, time zone, currency
    * Main domain
    * Flag URL
    * Domains and call codes
* JSON output for easy integration in scripts

## 📦 Installation

```bash
git clone https://github.com/extractip/ipwho.git
cd ipwho
go build -o ipwho
sudo mv ipwho /usr/local/bin/
```

Now the utility is available globally as `ipwho`.

## 🖥 Usage Examples

* Show only the country:

  ```bash
  ipwho 8.8.8.8
  ```

  ```
  United States
  ```

* Verbose output:

  ```bash
  ipwho -v 8.8.8.8
  ```

  ```
  Exit IP:       8.8.8.8
  Country:       United States (US)
  Region:        California
  Capital:       Washington
  Time Zone:     America/Los_Angeles
  Currency:      USD
  Domain:        google.com
  Flag URL:      https://api.extractip.com/flags/3x2/US.svg
  Domains:       google.com, youtube.com
  Call Codes:    +1
  ```

* JSON output:

  ```bash
  ipwho --json 8.8.8.8
  ```

  ```json
  {
    "exit_ip": "8.8.8.8",
    "country_name": "United States",
    "country_code": "US",
    "country_region": "California",
    "capital": "Washington",
    "time_zone": "America/Los_Angeles",
    "currency": "USD",
    "domain_name": "google.com",
    "flag_url": "[US flag](https://api.extractip.com/flags/3x2/US.svg)",
    "domains": ["google.com", "youtube.com"],
    "call_codes": ["+1"]
  }
  ```

## ⚙️ Flags

* `-v, --verbose` — show extended details in a table
* `--json` — output raw JSON instead of table

## 📜 License

MIT License.
