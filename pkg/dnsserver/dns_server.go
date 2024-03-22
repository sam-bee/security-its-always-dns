package dnsserver

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/miekg/dns"
	"github.com/sam-bee/security-itsalwaysdns/pkg/persistence"
	"log"
	"net"
	"strings"
)

var ipAddress string
var dnsPort string
var database *persistence.Database

type handler struct{}

func RunNameserver(ipAddressToReturn string, dnsPortNumber string, db *persistence.Database) {
	ipAddress = ipAddressToReturn
	dnsPort = dnsPortNumber
	database = db
	err := db.Initialise()
	if err != nil {
		log.Fatalf("Failed to initialise database: %s\n", err.Error())
	}
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
		domainToPersist := strings.TrimRight(domain, ".")
		persist(domainToPersist)

		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
			A:   net.ParseIP(ipAddress),
		})
	}
	w.WriteMsg(&msg)
}

func persist(fqdn string) {
	err := database.Store(fqdn)

	if err != nil {
		log.Fatalf("Failed to store query for %s: %s\n", fqdn, err.Error())
	} else {
		fmt.Printf("Received and stored query for %s\n", fqdn)
	}
}
