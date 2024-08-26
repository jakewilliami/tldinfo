package tldinfo

import "testing"

func TestTLDMap(t *testing.T) {
	domain := ".aaa"
	aaaTld := TLDInfoMap[domain]
	if aaaTld.Domain != domain {
		t.Fatalf(`Expected %q, found %q`, domain, aaaTld.Domain)
	}
}
