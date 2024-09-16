package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// redisConfig redis configuration
type redisConfig struct {
	Addr string `yaml:"addr"`
	DB   int    `yaml:"db"`
}

// Config configuration
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
		Kind  string      `yaml:"kind"`
		Redis redisConfig `yaml:"redis"`
	} `yaml:"pubsub"`
	Storage struct {
		Kind  string      `yaml:"kind"`
		Redis redisConfig `yaml:"redis"`
	} `yaml:"storage"`
}

// FromFile init config from file
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
