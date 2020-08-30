package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GetConfig(configFilePath string) Config {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal("Could not read config file: ", err)
	}
	conf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	if err := json.Unmarshal(conf, &config); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Configuration found ")
	return config
}
