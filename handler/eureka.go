package handler

import (
	"github.com/ddatsh/go-eureka-server/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var appMapMutex = sync.Mutex{}
var appMap = map[string]map[string]model.InstanceInfo{}

func Instance(c *gin.Context) {

	serviceName := c.Param("serviceName")
	var json model.Instance
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, ok := appMap[serviceName]
	if !ok {
		println("add ", serviceName)
		instanceMap := map[string]model.InstanceInfo{}
		instanceMap[json.Instance.InstanceID] = json.Instance
		appMapMutex.Lock()
		appMap[c.Param("serviceName")] = instanceMap
		appMapMutex.Unlock()
	} else {
		instanceMap := appMap[serviceName]
		_, instanceExists := instanceMap[json.Instance.InstanceID]
		if !instanceExists {
			appMapMutex.Lock()
			instanceMap[json.Instance.InstanceID] = json.Instance
			appMapMutex.Unlock()
		}
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

func getApps() model.Applications {
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
	return deltaInfo
}
func Apps(c *gin.Context) {

	deltaInfo := getApps()

	c.XML(200, deltaInfo)
}

func Info(c *gin.Context) {
	m := gin.H{
	}
	deltaInfo := getApps()
	for _, v := range deltaInfo.Applications {
		instanceInfo := make([]string, 0, len(v.Instances))
		for _, instance := range v.Instances {
			instanceInfo = append(instanceInfo, instance.HostName+":"+strconv.Itoa(instance.Port.Port))
		}
		m[v.Name] = strings.Join(instanceInfo, ",")
	}

	c.JSON(200, m)
}
