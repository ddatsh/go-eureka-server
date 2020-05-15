package route

import (
	"github.com/ddatsh/go-eureka-server/handler"
	"github.com/gin-gonic/gin"
	"log"
)


func EurekaInit(port string) {

	r := gin.Default()
	r.RedirectTrailingSlash = true

	r.Any("", handler.Index)
	r.Static("/static", "./static")
	r.POST("/eureka/apps/:serviceName", handler.Instance)
	r.DELETE("/eureka/apps/:serviceName/:instance", handler.InstanceDelete)
	r.PUT("/eureka/apps/:serviceName/:instance", handler.InstancePut)
	r.GET("/eureka/apps/delta", handler.Apps)
	r.POST("/eureka/peerreplication/batch/", handler.PeerReplicationBatch)
	r.GET("/eureka/apps/", handler.Apps)
	r.GET("/eureka/info", handler.Info)

	log.Fatal(r.Run(":" + port))
}
