package main

import (
	"github.com/gin-gonic/gin"
)

/*
	将url中传入的的信息,抽象到一个结构中,然后操作
*/

// 声明一个结构体,包含前端提供的字段
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 用于接收 url 中传递过来的参数
	router.GET("/login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindQuery(&json); err == nil {
			if json.User == "abc" && json.Password == "123" {
				c.JSON(200, gin.H{"message": "welcome"})
			} else {
				c.JSON(401, gin.H{"message": "wrong username or password"})
			}
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})

	router.Run()
}
