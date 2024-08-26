We want to check which ccTLDs have no associated country information (so that we can investigate why).

```bash
$ go run tools/cctldmissing/cctldmissing.go
[INFO] found table [columns [Domain Type TLD Manager] count 1591]
[INFO] found table [columns [Domain Names Root Zone Registry.INT Registry.ARPA RegistryIDN Repository] count 3]
[WARNING] Count not find country information associated with ccTLD ".ac"
[WARNING] Count not find country information associated with ccTLD ".eu"
[WARNING] Count not find country information associated with ccTLD ".su"
```

Explaination of the above ccTLDs' missing countries: see [here](https://github.com/biter777/countries/issues/64#issuecomment-1803391519) and [here](https://github.com/jakewilliami/tldinfo/blob/0a70e06eb23e9db03b4669ed3edcfd03ed4f27f8/tools/writetlds/writetlds.go#L112-L127).  Saint Helena has a ccTLD .ac, but country code SH.  .su is for Soviet Union, but as it is no longer a country (e.g., ISO 3166-3).  .eu is the country code for the European Union, which isn't a country.
