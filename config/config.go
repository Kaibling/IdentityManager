package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"iotuil"
	"os"
)

var Configuration *Config = nil
var configFilePath = ".im.json"

type Config struct {
	Email      string
	DBFilePath string
}

func InitConfig() {
	configData, err := readConfigFile()
	if err != nil {
		panic(err)
	}
	Configuration = configData
	return
}

func defaultConfig() *Config {
	return &Config{Email: "change@me.now", DBFilePath: "im.csv"}
}

func readConfigFile() (*Config, error) {
	var f *os.File
	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		f, err = os.Create(configFilePath)
		if err != nil {
			return nil, err
		}
		f.Close()

	}
	body, err := iotuil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	var c Config
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
