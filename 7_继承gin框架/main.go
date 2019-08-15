package main

/*
安装 gin
go get -u github.com/gin-gonic/gin
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"message": "pong",
		//})
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000") // 默认开启的端口为 8000
}
