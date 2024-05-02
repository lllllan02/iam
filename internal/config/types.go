package config

import "golang.org/x/oauth2/jwt"

type Config struct {
	Mysql Mysql      `yaml:"mysql"`
	Jwt   jwt.Config `yaml:"jwt"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
