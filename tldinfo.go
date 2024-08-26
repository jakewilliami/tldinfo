package tldinfo

// Generate new const library file with go generate
// Idea from Simon Sawert:
// https://github.com/bombsimon/tld-validator/blob/c0d0fbf9/tld.go#L9
//go:generate echo "[INFO] Generating library file"
//go:generate go run tools/writetlds/writetlds.go const
//go:generate echo "[INFO] Generating JSON file"
//go:generate go run tools/writetlds/writetlds.go json

// TLD types
// https://stackoverflow.com/a/71934535/
type TLDType string

const (
	Generic           TLDType = "generic"
	CountryCode       TLDType = "country-code"
	Sponsored         TLDType = "sponsored"
	Test              TLDType = "test"
	GenericRestricted TLDType = "generic-restricted"
	Infrastructure    TLDType = "infrastructure"
)

type TLD struct {
	Domain  string
	Type    TLDType
	Manager string
	Country string
}
