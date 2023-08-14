package config

import (
	"niltasks-auth/pkg/postgres/config"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

type Config struct {
	Server ServerConfig          `yaml:"server"`
	Mongo  config.PostgresConfig `yaml:"postgres"`
}

func MustLoad() *Config {
	var cfg Config

	err := cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
