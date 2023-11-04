package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/jakewilliami/tldeets/pkg/tldeets"
)

type TLDTypeCount struct {
	Type  tldeets.TLDType
	Count int
}

// https://stackoverflow.com/a/38644571
var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	rootpath   = filepath.Dir(filepath.Dir(basepath))
)

func main() {
	dataPath := filepath.Join(rootpath, "assets", "tlds.json")
	file, err := os.Open(dataPath)
	if err != nil {
		fmt.Printf("[ERROR] Could not read file \"%s\": %s\n", dataPath, err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	tlds := make(map[string]tldeets.TLD)
	err = decoder.Decode(&tlds)

	if err != nil {
		fmt.Printf("[ERROR] Could not deserialised JSON data from file \"%s\": %s\n", dataPath, err)
		os.Exit(1)
	}

	countMap := make(map[tldeets.TLDType]int)
	for _, v := range tlds {
		countMap[v.Type]++
	}

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

	for _, tc := range typeCounts {
		fmt.Printf("%-*s %*d\n", mtw+1, tc.Type, mcw+1, tc.Count)
	}
}
