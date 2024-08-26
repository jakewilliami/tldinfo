Helper tool to persist TLD data to disk.

Will pull from the internet and write to the appropriate directory.

Usage:
```bash
$ go run tools/writetlds/writetlds.go json
[INFO] Found base module path at ~/projects/tldinfo
[INFO] found table [columns [Domain Type TLD Manager] count 1591]
[INFO] found table [columns [Domain Names Root Zone Registry.INT Registry.ARPA RegistryIDN Repository] count 3]
[INFO] Wrote 210804 bytes to "~/projects/tldinfo/data/tlds.json"

$ go run tools/writetlds/writetlds.go const
[INFO] Found base module path at ~/projects/tldinfo
[INFO] found table [columns [Domain Type TLD Manager] count 1591]
[INFO] found table [columns [Domain Names Root Zone Registry.INT Registry.ARPA RegistryIDN Repository] count 3]
[WARNING] Skipped 169 non-unicode domain name(s)
[INFO] Wrote 233763 bytes to "~/projects/tldinfo/tldconsts.go"
[INFO] Successfully ran `go fmt` on output file "~/projects/tldinfo/tldconsts.go"
```
