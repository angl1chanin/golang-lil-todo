package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	App         string `yaml:"app"`
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func New() (*Config, error) {
	configPath := "./config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
		return nil, err
	}

	var cfg Config

	yamlCfg, _ := os.ReadFile(configPath)
	err := yaml.Unmarshal(yamlCfg, &cfg)
	if err != nil {
		log.Fatalf("Error while reading config file: %s", configPath)
		return nil, err
	}

	return &cfg, nil
}
