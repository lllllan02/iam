package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env string `default:"local"`

	Server struct {
		Host     string `default:"0.0.0.0"`
		HttpPort uint16 `default:"8000"`
		GrpcPort uint16 `default:"8001"`
	}

	Data struct {
		Mysql struct {
			Addr     string `default:"127.0.0.1:3306"`
			User     string `default:"root"`
			Password string `default:"123456"`
			DBName   string `default:"mysql"`
		}

		Redis struct {
			DB       int    `default:"0"`
			Addr     string `default:"127.0.0.1:6379"`
			Password string `default:""`
			Prefix   string `default:""`
		}
	}

	Log struct {
		Compress   bool   `default:"true"`              // Compression or not
		MaxAge     int    `default:"7"`                 // Maximum number of days the file can be saved
		MaxSize    int    `default:"1024"`              // Maximum size unit for each log file: M
		MaxBackups int    `default:"30"`                // The maximum number of backups that can be saved for log files
		Level      string `default:"info"`              // log level: debug, info, warn, error
		FileName   string `default:"./logs/server.log"` // log file path
		Encoding   string `default:"console"`           // json, console
	}
}

func NewConfig(path string) *Config {
	v := viper.New()

	initDefaults(v)

	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}
	return &conf
}

func initDefaults(v *viper.Viper) {
	initDefaultsRecursive(v, reflect.TypeOf(Config{}), "")
}

func initDefaultsRecursive(v *viper.Viper, t reflect.Type, prefix string) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ft := f.Type

		if f.Anonymous {
			initDefaultsRecursive(v, ft, prefix)
		} else {
			name := strings.ToLower(f.Name)
			if value, ok := f.Tag.Lookup("mapstructure"); ok {
				name = value
			}
			if ft.Kind() == reflect.Struct {
				initDefaultsRecursive(v, ft, fmt.Sprintf("%s%s.", prefix, name))
			} else {
				if value, ok := f.Tag.Lookup("default"); ok {
					key := prefix + name
					v.SetDefault(key, value)
				}
			}
		}
	}
}

func (conf *Config) IsProdEnv() bool {
	return strings.Contains(conf.Env, "prod")
}
