package main

import (
	"github.com/ddatsh/go-eureka-server/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func ConfigServerInit(port string) {

	r := gin.Default()
	r.RedirectTrailingSlash = true

	r.Any("", func(c *gin.Context) {
		c.String(200, "config server running")
	})

	r.GET("/:name", handler.GetConfig)
	r.GET("/:name/:env", handler.GetConfigWithEnv)

	log.Fatal(r.Run(":" + port))
	//log.Fatal(r.Run(":8887" ))
}
