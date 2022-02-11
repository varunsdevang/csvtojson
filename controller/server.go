package controller

import (
	echo "github.com/labstack/echo/v4"

	"github.com/swaggo/echo-swagger"

	_ "github/csvtojson/docs"
)

type ServerConfig struct {
	Host string
	Port string
}

func StartHTTPServer(config *ServerConfig) {
	e := echo.New()

	e.GET("api/swagger/*", echoSwagger.WrapHandler)

	h := NewConverterController()
	e.POST("api/v1/csvtojson", h.CsvToJsonHandler)

	e.Start(config.Host + ":" + config.Port)
}
