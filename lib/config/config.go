package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var Configuration *Config = nil
var configFilePath = ".im.json"

var dialects = map[string]int{"CSV": 1, "SQLITE": 1}

type Config struct {
	Email      string
	DBFilePath string
	Dialect    string
}

func InitConfig() error {
	configData, err := readConfigFile()
	if err != nil {
		return err
	}
	err = validate(configData)
	if err != nil {
		return err
	}
	Configuration = configData
	return nil
}
func validate(conf *Config) error {
	if _, ok := dialects[conf.Dialect]; !ok {
		return fmt.Errorf("unknown dialect '%s'", conf.Dialect)
	}
	return nil
}

func defaultConfig() Config {
	return Config{Email: "@change.me", DBFilePath: ".im.db", Dialect: "SQLITE"}
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
