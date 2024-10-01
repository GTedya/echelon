package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	GrpcPort string `yaml:"grpc_port"`
	DBPath   string `yaml:"database_path"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("file reading: %w", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("yaml unmarshalling: %w", err)
	}

	return &config, nil
}
