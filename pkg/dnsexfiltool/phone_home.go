package dnsexfiltool

func PhoneHome(mainDomain string) {
	exfilPayload := "hello world"
	dns := DnsLookupTool{}
	exfilData(exfilPayload, mainDomain, &dns)
}
