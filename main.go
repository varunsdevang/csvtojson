package main

import (
	"github/csvtojson/config"
	"github/csvtojson/controller"

	log "github.com/golang/glog"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Info("error while loading app config...", err.Error())
		return
	}

	controller.StartHTTPServer(&controller.ServerConfig{Host: config.HTTPHost, Port: config.HTTPPort})
}
