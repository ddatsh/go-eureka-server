package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var EurekaEnabled = false
var ConfigServerEnabled = false
var EurekaPort = "8761"
var ConfigServerPort = "8888"

func main() {

	flag.BoolVar(&EurekaEnabled, "eurekaEnabled", true, "enable eureka server")
	flag.StringVar(&EurekaPort, "eurekaPort", "8761", "eureka port")
	flag.BoolVar(&ConfigServerEnabled, "configServer", true, "enable config server")
	flag.StringVar(&ConfigServerPort, "configServerPort", "8888", "config server port")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	if EurekaEnabled {
		go func() {
			EurekaInit(EurekaPort)
		}()
	}
	if ConfigServerEnabled {
		go func() {
			ConfigServerInit(ConfigServerPort)
		}()
	}

	select {}

}
