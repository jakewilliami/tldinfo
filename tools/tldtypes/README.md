We want to check what TLD types we have (and check that our list is complete inside [the `tldinfo` library](../../tldinfo.go).

Will pull from the internet and print a sorted list to stdout.  Need to check against `TLDType` manually.

```bash
 $ go run tools/tldtypes/tldtypes.go
[INFO] found table [columns [Domain Type TLD Manager] count 1591]
[INFO] found table [columns [Domain Names Root Zone Registry.INT Registry.ARPA RegistryIDN Repository] count 3]
[INFO] Found 6 unique TLD types from 1591

TLD type frequencies:
=========================
generic              1246
country-code          316
sponsored              14
test                   11
generic-restricted      3
infrastructure          1
```
