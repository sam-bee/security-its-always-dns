package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFileFlag string

var rootCmd = &cobra.Command{
	Use:   "dnsreceiver",
	Short: "DNS Receiver component of the itsalwaysdns project.",
	Long: `Listens to data being exfiltrated in the form of DNS lookups. Stores incoming data in an SQLite database.
Can be queried for data.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFileFlag, "config", "", "config file (default is ./.env)")
}

func initConfig() {
	readConfigFile()
	readEnvVars()
}

func readConfigFile() {
	if configFileFlag != "" {
		viper.SetConfigFile(configFileFlag)
	} else {
		pwd, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(pwd)
		viper.SetConfigType("env")
		viper.SetConfigName(".env")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func readEnvVars() {
	viper.AutomaticEnv()
}
