package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	// https://stackoverflow.com/a/74328802
	"github.com/nfx/go-htmltable"
)

// TLD types
// https://www.sohamkamani.com/golang/enums/
const (
	Generic     string = "generic"
	CountryCode        = "country-code"
)

type TLD struct {
	Domain  string `header:"Domain"`
	Type    string `header:"Type"`
	Manager string `header:"TLD Manager"`
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

	data := make(map[string]TLD, len(table))
	for i := 0; i < len(table); i++ {
		tld := table[i]
		data[tld.Domain] = tld
	}

	tldJson, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("[ERROR] Could not JSONify data: %s\n", err)
		os.Exit(1)
	}

	outFile := "../../assets/tlds.json"
	err = ioutil.WriteFile(outFile, tldJson, 0644)
	if err != nil {
		fmt.Printf("[ERROR] Count not write JSON output to %s: %s", outFile, err)
		os.Exit(1)
	}

	// fmt.Printf("[DEBUG] %+v\n", data)
	fmt.Printf("[INFO] Wrote %d bytes to \"%s\"\n", len(tldJson), outFile)
}
