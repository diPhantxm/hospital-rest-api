package storage

import "fmt"

type Config struct {
	DataSourceName string `toml:"data_source_name"`
}

func NewConfig() *Config {
	return &Config{
		DataSourceName: "",
	}
}

func (c *Config) Print() {
	fmt.Println(c.DataSourceName)
}
