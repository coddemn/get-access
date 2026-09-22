package config

import (
	"os"
	"time"

	yaml "go.yaml.in/yaml/v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DBConfig struct {
	Host            string        `yaml:"host"`
	Port            string        `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Name            string        `yaml:"name"`
	SSLMode         string        `yaml:"sslmode"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if v := os.Getenv("DB_HOST"); v != "" {
		cfg.DB.Host = v
	}

	if v := os.Getenv("DB_PORT"); v != "" {
		cfg.DB.Port = v
	}

	if v := os.Getenv("DB_USER"); v != "" {
		cfg.DB.User = v
	}

	if v := os.Getenv("DB_NAME"); v != "" {
		cfg.DB.Name = v
	}

	if v := os.Getenv("DB_PASSWORD"); v != "" {
		cfg.DB.Password = v
	}

	if v := os.Getenv("DB_SSLMODE"); v != "" {
		cfg.DB.SSLMode = v
	}

	return &cfg, nil
}
