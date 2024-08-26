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

## Similar Projects

This library/CLI is **not** a top-level domain parser.  For parsers, see [`github.com/jpillora/go-tld`](https://github.com/jpillora/go-tld)/[`golang.org/x/net/publicsuffix`](https://pkg.go.dev/golang.org/x/net/publicsuffix)/[`github.com/bobesa/go-domain-util`](https://github.com/bobesa/go-domain-util).  It is more similar to a top-level domain *validator* such as [`github.com/bombsimon/tld-validator`](https://github.com/bombsimon/tld-validator), which looks like a very good project, even with unicode support using the [`golang.org/x/net` IDNA submodule](https://pkg.go.dev/golang.org/x/net/idna).  Simon's TLD validator even generates a constant/static list of TLDs for library/offline use.  However, this package does not retain any of the other TLD information [defined by the IANA](https://www.iana.org/domains/root/db).

## Caveats

  - Currently only supports country code TLDs (ccTLDs), as this is my personal main usecase for the application.  It does not yet generic TLDs (gTLD).
  - Internationalised ccTLDs are not yet supported.
