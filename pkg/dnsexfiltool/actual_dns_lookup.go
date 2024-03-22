package dnsexfiltool

import (
	"net"
)

type DnsLookupTool struct{}

type DnsLookup interface {
	Lookup(domain string) ([]string, error)
}

var _ DnsLookup = (*DnsLookupTool)(nil)

func (d *DnsLookupTool) Lookup(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	return convertIps(ips), nil
}

func convertIps(ips []net.IP) []string {
	ipsStr := make([]string, len(ips))
	for _, ip := range ips {
		ipsStr = append(ipsStr, ip.String())
	}
	return ipsStr
}
