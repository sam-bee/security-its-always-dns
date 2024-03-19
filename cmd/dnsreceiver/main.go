package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sam-bee/security-itsalwaysdns/pkg/dnsserver"
	"github.com/sam-bee/security-itsalwaysdns/pkg/persistence"
	"os"
)

var configFile string
var ipAddressToReturn string
var dnsPortNumber string
var sqlitePath string

func main() {
	readFlags()
	loadConfig(&configFile)
	db := persistence.GetDb(sqlitePath)
	dnsserver.RunNameserver(ipAddressToReturn, dnsPortNumber, db)
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
	ipAddressToReturn = getSetting("ITSALWAYSDNS_IP_ADDRESS")
	dnsPortNumber = getSetting("ITSALWAYSDNS_DNS_PORT_NUMBER")
	sqlitePath = getSetting("ITSALWAYSDNS_SQLITE_PATH")
}

func getSetting(setting string) string {
	value := os.Getenv(setting)
	if value == "" {
		panic(fmt.Sprintf("%s cannot be empty. Please set in the .env file or an environment variable.", setting))
	}
	return value
}
