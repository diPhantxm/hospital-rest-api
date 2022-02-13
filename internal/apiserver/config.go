package apiserver

import (
	"github.com/diphantxm/hospital-rest-api/internal/storage"
)

type Config struct {
	Port          string          `toml:"port"`
	Address       string          `toml:"address"`
	LogLevel      string          `toml:"log_level"`
	StorageConfig *storage.Config `toml:"storage"`
}

func NewConfig() *Config {
	return &Config{
		Port:          "",
		Address:       "",
		LogLevel:      "debug",
		StorageConfig: storage.NewConfig(),
	}
}
