package route

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"runtime"

	"github.com/ddatsh/go-eureka-server/handler"
	"github.com/gin-gonic/gin"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

type Person struct {
	Name    string
	Age     int
	Emails  []string
	Company string
	Role    string
}

type OnlineUser struct {
	User      []*Person
	LoginTime string
}

func EurekaInit(port string) {

	r := gin.Default()
	r.RedirectTrailingSlash = true

	r.Any("", func(c *gin.Context) {

		m := gin.H{
			"instanceInfo": gin.H{"ipAddr": "127.0.0.1",
				"status": "UP"},
			"generalInfo": gin.H{"total-avail-memory": "10MB",
				"environment":   "dev",
				"num-of-cpus":   runtime.NumCPU(),
				"server-uptime":"1 min"},
		}

		t, err := template.ParseFiles("tmpl.html")
		checkError(err)

		err = t.Execute(c.Writer, m)
		checkError(err)

		// c.String(200, "eureka server running")
	})

	r.Static("/static", "./static")
	r.POST("/eureka/apps/:serviceName", handler.Instance)
	r.DELETE("/eureka/apps/:serviceName/:instance", handler.InstanceDelete)
	r.PUT("/eureka/apps/:serviceName/:instance", handler.InstancePut)
	r.GET("/eureka/apps/delta", handler.Apps)
	r.GET("/eureka/apps/", handler.Apps)
	r.GET("/eureka/info", handler.Info)

	log.Fatal(r.Run(":" + port))
}
