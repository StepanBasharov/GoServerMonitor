package tools

import (
	"github.com/pelletier/go-toml"
	"log"
	"servermonitor/pkg/schemas"
)

func ConfigReader(configPath string) (*schemas.Config, error) {
	config, err := toml.LoadFile(configPath)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	var conf schemas.Config

	err = config.Unmarshal(&conf)

	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return &conf, nil
}
