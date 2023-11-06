<h1 align="center">TLD Info</h1>

A small command-line tool to provide basic information about a top-level domain (TLD).

---

## Quick Start

### Using the Library

```bash
$ go get github.com/jakewilliami/tldinfo
```

```go
package main

import (
	"fmt"

	"github.com/jakewilliami/tldinfo"
)

func main() {
	fmt.Printf("%v\n", tldinfo.TLDInfoMap[".jp"].Country)  // Japan
}
```

### Compiling the CLI

```bash
$ ./build.sh    # All-in-one build script
$ go build -o ./tldinfo cmd/tldinfo/main.go  # Or build it into a binary
```

## Caveats

  - Currently only supports country code TLDs (ccTLDs), as this is my personal main usecase for the application.  It does not yet generic TLDs (gTLD).
  - Internationalised ccTLDs are not yet supported.
