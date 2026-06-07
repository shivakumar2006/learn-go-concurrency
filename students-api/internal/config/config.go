package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	DatabaseURL string `yaml:"database_url" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var ConfigPath string

	ConfigPath = os.Getenv("CONFIG_PATH")

	if ConfigPath == "" {
		flags := flag.String("config", "", "Path to config file")
		flag.Parse()

		ConfigPath = *flags

		if ConfigPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesn't exist %s", ConfigPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(ConfigPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file %s", err.Error())
	}

	return &cfg
}
