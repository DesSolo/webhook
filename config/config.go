package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BindAddress string `yaml:"bind_address"`
	Logging     struct {
		Handler string `yaml:"handler"`
		Option  struct {
			AddSource bool `yaml:"add_source"`
			Level     int  `yaml:"level"`
		} `yaml:"options"`
		Args map[string]string `yaml:"args"`
	} `yaml:"logging"`
	Server struct {
		ServeStatic bool   `yaml:"serve_static"`
		StaticPath  string `yaml:"static_path"`
	}
	PubSub struct {
		Kind  string `yaml:"kind"`
		Redis struct {
			Addr string `yaml:"addr"`
			DB   int    `yaml:"db"`
		} `yaml:"redis"`
	} `yaml:"pubsub"`
}

// FromFile инициализирует конфигурацию из файла
func FromFile(filePath string) (*Config, error) {
	body, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(body, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
