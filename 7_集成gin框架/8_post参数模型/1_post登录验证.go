package main

import (
	"github.com/gin-gonic/gin"
)

/*
	获取表单中的数据进行验证,返回相应的信息
*/
func main1() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		// 获取表单中提交的数据
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 判断用户名和密码是否正确
		if username == "abc" && password == "123" {
			c.JSON(200, gin.H{"message": "welcome"})
		} else {
			c.JSON(401, gin.H{"message": "wrong username or password"})
		}
	})

	router.Run()
}
