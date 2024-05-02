package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func NewConfig() *Config {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	fmt.Printf("config: %+v\n", config)

	return &config
}
