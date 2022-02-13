package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/diphantxm/hospital-rest-api/internal/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	apiserver := apiserver.NewApiServer(config)

	apiserver.Configure()
	apiserver.Run()

	apiserver.Stop()
}
