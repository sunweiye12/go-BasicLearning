package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	将表单中传入的的信息,抽象到一个结构中,然后操作
*/
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 这种方法用于接收前端传递的 json 数据
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err == nil {
			if json.User == "abc" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			println("========出错啦=======")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 这种方案用于处理前端传递的 from 表单
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login

		if err := c.ShouldBind(&form); err == nil {
			if form.User == "abc" && form.Password == "123" {
				// 验证正确
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				// 密码过用户名不正确
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.Run()
}
