package main

import (
	"github.com/gin-gonic/gin"
)

/*
gin框架封装了http库，提供了GET、POST、PUT、DELETE、PATCH、HEAD、OPTIONS这些http请求方式。
使用router.method()来绑定路由
** 注意: 使用对应的method访问对应的path，会的对应的response。
  但是如果使用不对应的method访问path，就会返回404，比如使用post方法访问localhost:8080/get会返回404
*/
func main() {
	router := gin.Default()
	router.GET("/get", func(c *gin.Context) { c.JSON(200, gin.H{"message": "use get method"}) })
	router.POST("/post", func(c *gin.Context) { c.JSON(200, gin.H{"message": "use post method"}) })
	router.PUT("/put", func(c *gin.Context) { c.JSON(200, gin.H{"message": "use put method"}) })
	router.DELETE("/delete", func(c *gin.Context) { c.JSON(200, gin.H{"message": "use delete method"}) })
	router.PATCH("/patch", func(c *gin.Context) { c.JSON(200, gin.H{"message": "use patch method"}) })
	router.HEAD("/head", func(c *gin.Context) { c.JSON(200, gin.H{"message": "use head method"}) })
	router.OPTIONS("/options", func(c *gin.Context) { c.JSON(200, gin.H{"message": "use options method"}) })
	router.Run() // 应用默认绑定 8080 端口
}
