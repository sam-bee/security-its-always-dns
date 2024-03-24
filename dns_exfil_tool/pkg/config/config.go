package config

import (
	_ "embed"
	"github.com/pelletier/go-toml"
)

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

var configIsInitialised = false

func Initialise(configFileContents string) {
	err := toml.Unmarshal([]byte(configFileContents), &config)
	if err != nil {
		panic(err)
	}
	validateConfig()
	configIsInitialised = true
}

func IsPhoneHomeMissionEnabled() bool {
	checkConfigConfigured()
	return config.Missions.PhoneHome
}

func IsExfilFilesEnabled() bool {
	checkConfigConfigured()
	return config.Missions.ExfilFiles
}

func IsExfilEnvVarsEnabled() bool {
	checkConfigConfigured()
	return config.Missions.ExfilEnvVars
}

func GetFilesAndFoldersToExfil() []string {
	checkConfigConfigured()
	return config.Objectives.FilesAndFolders
}

func GetMainDomain() string {
	checkConfigConfigured()
	return config.ExfilServer.MainDomain
}

func GetPortNo() int {
	checkConfigConfigured()
	return config.ExfilServer.PortNo
}

func checkConfigConfigured() {
	if !configIsInitialised {
		panic("Config error: not initialised")
	}
}

func validateConfig() {
	if !config.Missions.PhoneHome && !config.Missions.ExfilFiles && !config.Missions.ExfilEnvVars {
		panic("Config error: no missions enabled")
	}
	if config.Missions.ExfilFiles && len(config.Objectives.FilesAndFolders) == 0 {
		panic("Config error: no files or folders to exfil")
	}
}
