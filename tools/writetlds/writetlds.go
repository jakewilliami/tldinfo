package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	// https://stackoverflow.com/a/74328802
	"github.com/nfx/go-htmltable"
	"github.com/biter777/countries"

	"github.com/jakewilliami/tldeets/pkg/tldeets"
)

// https://stackoverflow.com/a/38644571
var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	rootpath   = filepath.Dir(filepath.Dir(basepath))
)

type TLD struct {
	Domain  string          `header:"Domain"`
	Type    tldeets.TLDType `header:"Type"`
	Manager string          `header:"TLD Manager"`
}

func main() {
	fmt.Printf("[INFO] Found base module path at %s\n", rootpath)

	htmltable.Logger = func(_ context.Context, msg string, fields ...any) {
		fmt.Printf("[INFO] %s %v\n", msg, fields)
	}

	url := "https://www.iana.org/domains/root/db"
	table, err := htmltable.NewSliceFromURL[TLD](url)
	if err != nil {
		fmt.Printf("[ERROR] Could not get table by %s: %s", url, err)
		os.Exit(1)
	}

	dataRaw := make(map[string]TLD, len(table))
	for i := 0; i < len(table); i++ {
		tld := table[i]
		dataRaw[tld.Domain] = tld
	}

	data := make(map[string]tldeets.TLD, len(dataRaw))
	for tldStr, tld := range dataRaw {
		var country string
		// TODO: this will not always work; e.g. Saint Helena is has ccTLD .ac,
		// but country code SH.  Another example: .su is for Soviet Union, but
		// as it is no longer a country (e.g., ISO 3166-3).
		if tld.Type == tldeets.CountryCode {
			var countryCode string
			if tldStr[0] == '.' {
				countryCode = tldStr[1:]
			}
			countryCode = strings.ToUpper(countryCode)
			country = countries.ByName(countryCode).Info().Name
			if country == "Unknown" {
				country = ""
			}
		}
		data[tldStr] = tldeets.TLD{
			Domain: tld.Domain,
			Type: tld.Type,
			Manager: tld.Manager,
			Country: country,
		}
	}

	tldJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("[ERROR] Could not JSONify data: %s\n", err)
		os.Exit(1)
	}

	outFile := filepath.Join(rootpath, "assets", "tlds.json")
	err = ioutil.WriteFile(outFile, tldJson, 0644)
	if err != nil {
		fmt.Printf("[ERROR] Count not write JSON output to %s: %s", outFile, err)
		os.Exit(1)
	}

	// fmt.Printf("[DEBUG] %+v\n", data)
	fmt.Printf("[INFO] Wrote %d bytes to \"%s\"\n", len(tldJson), outFile)
}
