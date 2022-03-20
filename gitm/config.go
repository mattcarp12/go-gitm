package gitm

import (
	"encoding/json"
	"os"
)

type GitmConfig struct {
	Bare bool `json:"bare"`
}

// Read reads config from the .gitm/config file
func ReadConfig() GitmConfig {
	f := Files{}
	file, err := os.Open(f.GitmPath("config"))
	if err != nil {
		panic(err)
	}
	config := GitmConfig{}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}

// Write writes (overwrites) to the .gitm/config file
func WriteConfig(config GitmConfig) {
	f := Files{}
	file, err := os.Create(f.GitmPath("config"))
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(file).Encode(&config)
	if err != nil {
		panic(err)
	}
}

func IsBare() bool {
	config := ReadConfig()
	return config.Bare
}

func AssertNotBare() {
	if IsBare() {
		panic("this operation must be run in a work tree")
	}
}