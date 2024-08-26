package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jakewilliami/tldinfo"
)

// https://stackoverflow.com/a/38644571
var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	rootpath   = filepath.Dir(filepath.Dir(basepath))
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Must specify a TLD to search for")
		os.Exit(1)
	}

	// TODO: handle multiple arguments
	tldStr := os.Args[1]
	if tldStr[0] != '.' {
		tldStr = "." + tldStr
	}
	tld := tldinfo.TLDInfoMap[tldStr]
	country := tld.Country
	fmt.Printf("%v\n", country)
}

func _mainJSON() {
	dataPath := filepath.Join(rootpath, "data", "tlds.json")
	file, err := os.Open(dataPath)
	if err != nil {
		fmt.Printf("[ERROR] Could not read file \"%s\": %s\n", dataPath, err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	tlds := make(map[string]tldinfo.TLD)
	err = decoder.Decode(&tlds)

	if err != nil {
		fmt.Printf("[ERROR] Could not deserialised JSON data from file \"%s\": %s\n", dataPath, err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Must specify a TLD to search for")
		os.Exit(1)
	}

	// TODO: handle multiple arguments
	tldStr := os.Args[1]
	if tldStr[0] != '.' {
		tldStr = "." + tldStr
	}
	tld := tlds[tldStr]
	fmt.Println(tld)
}
