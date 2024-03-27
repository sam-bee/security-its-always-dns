package cmd

import (
	"github.com/sam-bee/security-itsalwaysdns/dns_receiver/pkg/nameserver"
	"github.com/sam-bee/security-itsalwaysdns/dns_receiver/pkg/persistence"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listens for incoming DNS lookups with data hidden inside",
	Long:  `Receives DNS lookups and stores the data hidden in subdomains in an SQLite database.`,
	Run: func(cmd *cobra.Command, args []string) {
		listenForDnsLookups()
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
}

func listenForDnsLookups() {

	ip := viper.GetString("itsalwaysdns_ip_address")
	port := viper.GetString("itsalwaysdns_dns_port_number")
	sqlite := viper.GetString("itsalwaysdns_sqlite_path")

	db := persistence.GetDb(sqlite)
	nameserver.RunNameserver(ip, port, db)
}
