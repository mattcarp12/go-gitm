package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/mattcarp12/go-gitm/gitm/files"
)

type GitmConfig struct {
	Bare bool `json:"bare"`
}

// Read reads config from the .gitm/config file
func ReadConfig() GitmConfig {
	file, err := os.Open(files.GitmPath("config"))
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
	file, err := os.Create(files.GitmPath("config"))
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
		log.Fatal("this operation must be run in a work tree")
	}
}