package main

import (
	"github.com/gin-gonic/gin"
)

/*
	分组路由是为了让程序看起来更加简洁清晰
	对于不同模块的访问采用不同的 URL,
	同一个模块为一组
*/
func main() {
	router := gin.Default()

	// 版本 1 访问 localhost:8080/v1/*
	v1 := router.Group("/v1")
	{
		v1.GET("/demo", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "/v1/demo"})
		})
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "/v1/test"})
		})
	}

	// 版本 2 访问 localhost:8080/v2/*
	v2 := router.Group("/v2")
	{
		v2.GET("/demo", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "/v2/demo"})
		})
		v2.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "/v2/test"})
		})
	}
	router.Run()
}
