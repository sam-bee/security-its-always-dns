package dnsserver

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/miekg/dns"
	"log"
	"net"
)

var ipAddress string
var dnsPort string

type handler struct{}

func RunNameserver(ipAddressToReturn string, dnsPortNumber string) {
	ipAddress = ipAddressToReturn
	dnsPort = dnsPortNumber
	srv := &dns.Server{Addr: ":" + dnsPort, Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}

func (h *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		doSomethingWithTheFqdn(domain)

		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
			A:   net.ParseIP(ipAddress),
		})
	}
	w.WriteMsg(&msg)
}

func doSomethingWithTheFqdn(fqdn string) {
	fmt.Printf("Received query for %s\n", fqdn)
}
