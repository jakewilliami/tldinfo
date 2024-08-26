<h1 align="center">TLD Info</h1>

A small command-line tool to provide basic information about a top-level domain (TLD).

---

## Quick Start

### Using the Library

Get the module:
```bash
$ go get github.com/jakewilliami/tldinfo
```

Basic library usage:
```go
package main

import (
	"fmt"

	"github.com/jakewilliami/tldinfo"
)

func main() {
	tld := tldinfo.TLDInfoMap[".jp"]
	country := tld.Country
	fmt.Printf("%v\n", country)  // Japan
}
```

Types:
```go
type TLD struct {
        Domain  string
        Type    TLDType
        Manager string
        Country string
}

const (
        Generic           TLDType = "generic"
        CountryCode       TLDType = "country-code"
        Sponsored         TLDType = "sponsored"
        Test              TLDType = "test"
        GenericRestricted TLDType = "generic-restricted"
        Infrastructure    TLDType = "infrastructure"
)
```

### Using the CLI

Compile:
```bash
$ ./build.sh    # All-in-one build script
$ go build -o ./tldinfo cmd/tldinfo/main.go  # Or build it into a binary
```

Run:
```bash
$ ./tldinfo jp  # or ./tldinfo .jp
$ ./tldinfo -h  # help command coming soon™!
```

## Caveats

  - Currently only supports country code TLDs (ccTLDs), as this is my personal main usecase for the application.  It does not yet generic TLDs (gTLD).
  - Internationalised ccTLDs are not yet supported.
