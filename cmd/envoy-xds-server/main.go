package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"fedx-mpc-xds-server/internal/app/service/v1"
	"fedx-mpc-xds-server/internal/pkg/utils/v1"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "configFile", "./config.yaml", "Config File")
}

func main() {
	flag.Parse()

	config := getXDSServerConfig(configFile)
	service.RunServer(config)
}

func getXDSServerConfig(configFile string) *utils.XDSServerConfig {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	var config utils.XDSServerConfig
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
