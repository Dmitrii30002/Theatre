package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server         ServerConfig         `yaml:"server"`
	TheatreService TheatreServiceConfig `yaml:"theatre_service"`
}

type ServerConfig struct {
	Address         string `yaml:"address"`
	ReadTimeout     string `yaml:"read_timeout"`
	WriteTimeout    string `yaml:"write_timeout"`
	IdleTimeout     string `yaml:"idle_timeout"`
	ShutdownTimeout string `yaml:"shutdown_timeout"`
}

type TheatreServiceConfig struct {
	Address string `yaml:"address"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config %q: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config %q: %w", path, err)
	}
	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) Validate() error {
	if c.Server.Address == "" {
		return fmt.Errorf("server.address is required")
	}
	if strings.TrimSpace(c.TheatreService.Address) == "" {
		return fmt.Errorf("theatre_service.address is required")
	}
	return nil
}

func ParseDuration(value string, fallback time.Duration) time.Duration {
	if value == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return parsed
}
