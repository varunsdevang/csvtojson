package main

import (
	"github/csvtojson/config"
	"github/csvtojson/controller"

	log "github.com/golang/glog"
)

// @title           CSV to JSON
// @version         1.0
// @description     a simple REST serverto convert csv file to json.
// @termsOfService  http://swagger.io/terms/

// @contact.name   T S Varun
// @contact.url    https://www.linkedin.com/in/tsvarun/
// @contact.email  varundevang@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3030
// @BasePath  /api/v1
func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Info("error while loading app config...", err.Error())
		return
	}

	controller.StartHTTPServer(&controller.ServerConfig{Host: config.HTTPHost, Port: config.HTTPPort})
}
