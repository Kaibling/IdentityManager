package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

func defaultConfig() Config {
	return Config{Email: "@change.me", DBFilePath: ".im.csv"}
}

func readConfigFile() (*Config, error) {
	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		c := defaultConfig()
		b, _ := json.Marshal(&c)
		ioutil.WriteFile(configFilePath, b, 0644)
		return &c, nil
	}
	body, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	var c Config
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
