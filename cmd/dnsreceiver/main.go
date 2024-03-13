package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/miekg/dns"
	"log"
	"net"
	"os"
)

var ipAddressToReturn string

type handler struct{}

func main() {
	loadConfig()
	runNameserver()
}

func loadConfig() {
	godotenv.Load(".env")
	ipAddressToReturn = os.Getenv("ITSALWAYSDNS_IP_ADDRESS")
}

func runNameserver() {
	srv := &dns.Server{Addr: ":53", Net: "udp"}
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
			A:   net.ParseIP(ipAddressToReturn),
		})
	}
	w.WriteMsg(&msg)
}

func doSomethingWithTheFqdn(fqdn string) {
	fmt.Printf("Received query for %s\n", fqdn)
}
