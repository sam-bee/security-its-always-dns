package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sam-bee/security-itsalwaysdns/pkg/dnsserver"
	"os"
)

var configFile string
var ipAddressToReturn string
var dnsPortNumber string

func main() {
	readFlags()
	loadConfig(&configFile)
	dnsserver.RunNameserver(ipAddressToReturn, dnsPortNumber)
}

func readFlags() {
	configFile = ""
	flag.StringVar(&configFile, "config", "./.env", "path to the config file")
	flag.Parse()
}

func loadConfig(file *string) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		panic(fmt.Sprintf("Config file does not exist: %s", configFile))
	}

	godotenv.Load(*file)
	ipAddressToReturn = os.Getenv("ITSALWAYSDNS_IP_ADDRESS")
	dnsPortNumber = os.Getenv("ITSALWAYSDNS_DNS_PORT_NUMBER")

	if ipAddressToReturn == "" {
		panic("IP address to return cannot be empty")
	}

	if dnsPortNumber == "" {
		panic("DNS port number cannot be empty")
	}
}
