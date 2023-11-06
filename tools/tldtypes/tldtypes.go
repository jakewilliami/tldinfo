package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	// https://stackoverflow.com/a/74328802
	"github.com/nfx/go-htmltable"

	"github.com/jakewilliami/tldinfo"
)

type TLDTypeCount struct {
	Type  tldinfo.TLDType
	Count int
}

type TLD struct {
	Domain  string          `header:"Domain"`
	Type    tldinfo.TLDType `header:"Type"`
	Manager string          `header:"TLD Manager"`
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

	countMap := make(map[tldinfo.TLDType]int, len(table))
	for i := 0; i < len(table); i++ {
		tld := table[i]
		countMap[tld.Type]++
	}
	fmt.Printf("[INFO] Found %d unique TLD types from %d\n", len(countMap), len(table))

	typeCounts := make([]TLDTypeCount, len(countMap))
	i := 0
	for t, n := range countMap {
		typeCounts[i] = TLDTypeCount{Type: t, Count: n}
		i++
	}

	sort.Slice(typeCounts, func(i, j int) bool {
		return typeCounts[i].Count > typeCounts[j].Count
	})

	mtw, mcw := 0, 0
	for _, tc := range typeCounts {
		if len(tc.Type) > mtw {
			mtw = len(tc.Type)
		}
		if len(fmt.Sprint(tc.Count)) > mcw {
			mcw = len(fmt.Sprint(tc.Count))
		}
	}

	fmt.Println("\nTLD type frequencies:")
	fmt.Println(strings.Repeat("=", mtw+mcw+3))
	for _, tc := range typeCounts {
		fmt.Printf("%-*s %*d\n", mtw+1, tc.Type, mcw+1, tc.Count)
	}
}
