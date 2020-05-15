package main

import (
	"flag"
	"github.com/ddatsh/go-eureka-server/route"
)

var EurekaEnabled = false
var ConfigServerEnabled = false
var EurekaPort string
var ConfigServerPort string

func main() {

	flag.BoolVar(&EurekaEnabled, "eurekaEnabled", true, "enable eureka server")
	flag.StringVar(&EurekaPort, "eurekaPort", "8762", "eureka port")
	flag.BoolVar(&ConfigServerEnabled, "configServer", true, "enable config server")
	flag.StringVar(&ConfigServerPort, "configServerPort", "8888", "config server port")
	flag.Parse()

	//gin.SetMode(gin.ReleaseMode)

	if EurekaEnabled {
		go func() {
			route.EurekaInit(EurekaPort)
		}()
	}
	if ConfigServerEnabled {
		go func() {
			route.ConfigServerInit(ConfigServerPort)
		}()
	}

	select {}

}
