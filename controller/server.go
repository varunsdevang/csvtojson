package controller

import (
	echo "github.com/labstack/echo/v4"
)

type ServerConfig struct {
	Host string
	Port string
}

func StartHTTPServer(config *ServerConfig) {
	e := echo.New()
	h := NewConverterController()
	e.POST("api/v1/csvtojson", h.CsvToJsonHandler)
	e.Start(config.Host + ":" + config.Port)
}
