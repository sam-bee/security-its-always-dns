package main

import (
	_ "embed"
	"fmt"
	"github.com/go-ini/ini"
)

//go:embed config.ini
var iniContents string

func main() {
	parseConfig()
}

func parseConfig() {
	// parse ini string
	cfg, err := ini.Load([]byte(iniContents))
	if err != nil {
		fmt.Printf("Failed to parse ini file: %v", err)
		return
	}

	// Access the values from the ini file
	section := cfg.Section("det")
	value := section.Key("DET_SETTING").String()
	fmt.Println(value)
}
