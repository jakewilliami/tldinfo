package main

import (
	"flag"
	"fmt"

	"github.com/jakewilliami/tldeets/pkg/tldeets"
)

func main() {
	name := flag.String("n", "Sailor", "The name to greet")
	flag.Parse()
	tldeets.GreetSailor()
	fmt.Printf("[cmd] Hallo, %s!\n", *name)
}
