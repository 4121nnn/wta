package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	Env    string `yaml:"env" env-default:"local"`
	Name   string `yaml:"name" env-default:"App"`
	Server `yaml:"server"`
}

type Server struct {
	Address     string        `yaml:"address" env-default:"localhost:8082"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

var (
	config *Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			configPath = "config/local.yaml"
			log.Println("Config path not passed. Using default config path: config/local.yaml")
		}

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			log.Fatalf("config file %s does not exist", configPath)
		}
		config = &Config{}
		if err := cleanenv.ReadConfig(configPath, config); err != nil {
			log.Fatalf("cannot read config: %s ", err)
		}
	})

	return config
}
