package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Read(filePath string) *Config {
	// Read config file
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &Config{}
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return &Config{}
	}

	return &config
}
