package handler

import (
	"github.com/ddatsh/go-eureka-server/util"
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
	"log"
	"os"

	"strings"
)

var PWD string

func init() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	PWD = dir
}

func GetConfigWithEnv(c *gin.Context) {

	path := c.Param("name")
	env := c.Param("env")

	var contentMap map[interface{}]interface{}
	contentMap = parse(PWD, path+"-"+env+".yml")

	c.JSON(200, gin.H{
		"name":     path,
		"profiles": []string{env},
		"label":    nil,
		"version":  nil,
		"state":    nil,
		"propertySources": []gin.H{
			{
				"name":   "classpath:/config/" + path + "-" + env + ".yml",
				"source": util.ConvertMap(util.FlatMap(contentMap))},
		},
	})

}
func GetConfig(c *gin.Context) {

	if "/favicon.ico" == c.Request.RequestURI {
		c.Status(204)
		return
	}

	path := c.Param("name")

	var contentMap map[interface{}]interface{}

	suffix := path[strings.LastIndex(path, ".")+1:]

	switch suffix {
	case "yml", "yaml":
		contentMap = parse(PWD, path)
		bs, err := yaml.Marshal(contentMap)
		if err != nil {
			panic(err)
		}
		c.String(200, string(bs))
	case "json":
		contentMap = parse(PWD, path)
		c.JSON(200, util.ConvertMap(contentMap))
	default:
		c.Status(404)
	}

}

func parse(filePath, app string) map[interface{}]interface{} {

	app = app[0:strings.LastIndex(app, ".")]

	var contentMap = map[interface{}]interface{}{}

	dir := filePath + "/config/" + app
	if _, err := os.Stat(dir); !os.IsNotExist(err) {

		mergo.Merge(&contentMap, util.ReadYml( dir+"/application.yml"),mergo.WithOverride)
		mergo.Merge(&contentMap, util.ReadYml(dir+"/"+app+".yml"),mergo.WithOverride)

	} else {
		baseApp := app[0:strings.LastIndex(app, "-")]
		dir = filePath + "/config/" + baseApp

		mergo.Merge(&contentMap, util.ReadYml( dir+"/application.yml"),mergo.WithOverride)
		mergo.Merge(&contentMap, util.ReadYml( dir+"/"+baseApp+".yml"),mergo.WithOverride)
		mergo.Merge(&contentMap, util.ReadYml( dir+"/"+app+".yml"),mergo.WithOverride)

	}

	return contentMap
}
