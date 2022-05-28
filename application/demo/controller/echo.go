package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/reatang/gostart/application/demo/models"
	"github.com/reatang/gostart/pkg/client/xhttp"
	"github.com/reatang/gostart/pkg/logger"
	"github.com/reatang/gostart/pkg/redis"
	"net/http"
)

func ActionEcho(c *gin.Context) {

	// 日志
	logger.Info("一条日志")

	// 数据库
	user := models.GetUserDao().First(10)

	// redis expiration参数在redis server 小于6.0版本写0，大于等于6.0版本写redis.KeepTTL
	redis.GetDB().Set(c, "gotest:00001", 123, 0)
	data, _ := redis.GetDB().Get(c, "gotest:00001").Int()

	// http client
	client := xhttp.New("http://127.0.0.1:8001")
	response, _ := client.R().Get("/curl")

	c.String(http.StatusOK, "OK：%s\n%d\n%s", user.Name, data, response.String())
}

func ActionCurl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"say": "Hello",
	})
}
