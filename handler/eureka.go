package handler

import (
	"github.com/ddatsh/go-eureka-server/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var appMap = map[string]map[string]model.InstanceInfo{}

func Instance(c *gin.Context) {
	var json model.Instance
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, ok := appMap[c.Param("serviceName")]
	if !ok {
		println("add ", c.Param("serviceName"))
		instanceMap := map[string]model.InstanceInfo{}
		instanceMap[json.Instance.InstanceID] = json.Instance
		appMap[c.Param("serviceName")] = instanceMap
	} else {
		instanceMap := appMap[c.Param("serviceName")]
		instanceMap[json.Instance.InstanceID] = json.Instance
	}

	c.Status(204)
}

func InstanceDelete(c *gin.Context) {

	log.Printf("delete instance %s %s\n", c.Param("serviceName"), c.Param("instance"))

	c.JSON(200, nil)
}

func InstancePut(c *gin.Context) {

	log.Printf("put instance %s %s\n", c.Param("serviceName"), c.Param("instance"))
	c.JSON(200, nil)
}

func Apps(c *gin.Context) {

	applications := make([]model.Application, 0, len(appMap))

	for i, v := range appMap {
		application := model.Application{Name: i}

		var instanceInfos = make([]model.InstanceInfo, 0, len(v))
		for _, v1 := range v {
			instanceInfos = append(instanceInfos, v1)

		}

		application.Instances = instanceInfos

		applications = append(applications, application)

	}
	deltaInfo := model.Applications{
		VersionsDelta: 1,
		AppsHashcode:  "UP_1_",
		Applications:  applications,
	}
	c.XML(200, deltaInfo)
}
