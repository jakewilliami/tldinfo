package tldinfo

// TLD types
// https://stackoverflow.com/a/71934535/12069968
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
