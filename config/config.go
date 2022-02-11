package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppName    string `json:"appName"`
	HideBanner bool   `json:"hideBanner"`
	HTTPHost   string `json:"httpHost"`
	HTTPPort   string `json:"httpPort"`
}

func LoadConfig() (*Config, error) {
	f, err := os.Open("config/config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	conf := new(Config)
	err = json.NewDecoder(f).Decode(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
