package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"time"
)

// $ cat settings.ini
//
// ArrisHost="75.55.145.214"
// ArrisPass="blah"
//

type Config struct {
	ArrisHost        string
	ArrisPass        string
	DeliveryInterval time.Duration
	DatabaseURL      string
	DatabaseUser     string
	DatabasePassword string
	DatabaseDatabase string
}

var configfile = "settings.ini"

func ReadConfig() Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
