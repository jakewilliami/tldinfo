package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	// https://stackoverflow.com/a/74328802
	"github.com/biter777/countries"
	"github.com/nfx/go-htmltable"

	"github.com/jakewilliami/tldinfo"
)

type TLD struct {
	Domain  string          `header:"Domain"`
	Type    tldinfo.TLDType `header:"Type"`
	Manager string          `header:"TLD Manager"`
}

func isASCII(s string) bool {
	for _, char := range s {
		if char > 127 {
			return false
		}
	}
	return true
}

func main() {
	htmltable.Logger = func(_ context.Context, msg string, fields ...any) {
		fmt.Printf("[INFO] %s %v\n", msg, fields)
	}

	url := "https://www.iana.org/domains/root/db"
	table, err := htmltable.NewSliceFromURL[TLD](url)
	if err != nil {
		fmt.Printf("[ERROR] Could not get table by %s: %s", url, err)
		os.Exit(1)
	}

	for i := 0; i < len(table); i++ {
		tld := table[i]
		if !isASCII(tld.Domain) {
			// Skip non-ascii domains for now - will add back later
			// TODO: consider handling non-ASCII domains with Punycode; see
			// github.com/bombsimon/tld-validator
			continue
		}

		// TODO: this will not always work; e.g. Saint Helena is has ccTLD .ac,
		// but country code SH.  Another example: .su is for Soviet Union, but
		// as it is no longer a country (e.g., ISO 3166-3).
		if tld.Type == tldinfo.CountryCode {
			var countryCode string
			if tld.Domain[0] == '.' {
				countryCode = tld.Domain[1:]
			}
			countryCode = strings.ToUpper(countryCode)
			if countries.ByName(countryCode).Info().Name == "Unknown" {
				fmt.Printf("[WARNING] Count not find country information associated with ccTLD \"%s\"\n", tld.Domain)
			}
		}
	}
}
