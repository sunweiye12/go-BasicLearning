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

//声明一个gin.HandlerFunc
func LoginCheck(c *gin.Context) {
	var json Login
	var err error
	if err := c.ShouldBindQuery(&json); err == nil { //尝试从get请求中获取参数
		if json.User == "abc" && json.Password == "123" {
			c.JSON(200, gin.H{"message": "login in success by Get method"})
		} else {
			c.JSON(200, gin.H{"message": "login in failed by Get method"})
		}
	} else if err := c.ShouldBind(&json); err == nil { //尝试从post请求中获取参数
		if json.User == "abc" && json.Password == "123" {
			c.JSON(200, gin.H{"message": "login in success by POST method"})
		} else {
			c.JSON(200, gin.H{"message": "login in failed by POST method"})
		}
	}
	if err != nil { //解析请求中的参数失败
		c.JSON(400, gin.H{"message": err.Error()})
	}
}

// 将事件处理程序提取出来
func main() {
	router := gin.Default()
	//对于GET和POST，都尝试使用同一个事件处理程序--->(将事件处理程序提取出来)
	router.GET("login", LoginCheck)
	router.POST("login", LoginCheck)
	router.Run()
}
