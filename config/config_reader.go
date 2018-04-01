package config

import (
	"encoding/json"
	//"fmt"
	"log"
	"os"
)

func GetConfigurations() Configuration {
	file, err := os.Open("config.json")
	defer file.Close()

	defaults := false

	if err != nil {
		log.Printf("There was an error opening config.json - %s\n", err)
		defaults = true
	}

	config := Configuration{}
	if defaults {
		useDefaultConfig(&config)
	} else {
		log.Println("Using config values")
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&config)

		if err != nil {
			log.Printf("There was an error parsing config.json - %s\n", err)
			useDefaultConfig(&config)
		}
	}
	return config
}

func useDefaultConfig(config *Configuration) {
	log.Println("Using default values")
	config.IpAddress = "127.0.0.1"
	config.Port = 6911
	config.ReadTimeout = -1
	config.WriteTimeout = -1
	config.HtmlBaseFolder = "html/"
}
