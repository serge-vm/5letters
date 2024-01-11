package config

import (
	"flag"
	"strings"
)

type APIConfig struct {
	Host    string
	Port    int
	BaseUrl string
}

func NewAPIConfig() APIConfig {
	var config APIConfig

	hostParam := flag.String("host", "", "Host to listen on")
	portParam := flag.Int("port", 22080, "Port to listen on")
	baseUrlParam := flag.String("base-url", "", "Base URL")
	flag.Parse()

	config.Host = *hostParam
	config.Port = *portParam
	if *baseUrlParam != "" {
		config.BaseUrl = "/" + strings.Trim(*baseUrlParam, "/")
	}

	return config
}
