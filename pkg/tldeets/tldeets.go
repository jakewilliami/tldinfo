package tldeets

// TLD types
// https://stackoverflow.com/a/71934535/12069968
type TLDType string

const (
	Generic     TLDType = "generic"
	CountryCode TLDType = "country-code"
)

type TLD struct {
	Domain  string
	Type    TLDType
	Manager string
}
