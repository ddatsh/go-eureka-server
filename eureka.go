package main

import (
	"github.com/ddatsh/go-eureka-server/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func EurekaInit(port string){

	r := gin.Default()
	r.RedirectTrailingSlash = true

	r.Any("", func(c *gin.Context) {
		c.String(200, "eureka server running")
	})
	r.POST("/eureka/apps/:serviceName", handler.Instance)
	r.DELETE("/eureka/apps/:serviceName/:instance", handler.InstanceDelete)
	r.PUT("/eureka/apps/:serviceName/:instance", handler.InstancePut)
	r.GET("/eureka/apps/delta", handler.Apps)
	r.GET("/eureka/apps/", handler.Apps)

	log.Fatal(r.Run(":"+port))
}
