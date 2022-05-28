package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/reatang/gostart/application/demo/controller"
)

func RegisterRouter() *gin.Engine {
	engine := gin.Default()
	
	engine.GET("/", controller.ActionEcho)
	engine.GET("/curl", controller.ActionCurl)
	
	return engine
}
