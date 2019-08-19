package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 接收表单数据  (multipart/urlencoded)  --> 用于post 的表单或者urlencoded 提交的数据
	/*
		router.POST("/post", func(c *gin.Context) {
			name := c.PostForm("name")		// 通过表单中的 key 获取传递放 value
			age := c.PostForm("age")
			sex := c.DefaultPostForm("sex", "male")	// 如果传递表单中包含则获取,否则得到默认值
			c.JSON(200, gin.H{"name": name, "age": age, "sex": sex})
		})
	*/

	// 同时接收 URL 中传递的数据和 post 表单中传输的数据
	// 注意，这里必须使用router.POST
	router.POST("/post", func(c *gin.Context) {
		// 获取表单中传递的数据
		name := c.PostForm("name")
		//name := c.BindJSON("name")
		age := c.PostForm("age")
		sex := c.DefaultPostForm("sex", "male")
		// 获取 URL 中传递的数据
		addr := c.Query("addr")
		hobby := c.DefaultQuery("hobby", "basketball")
		c.JSON(200, gin.H{"name": name, "age": age, "sex": sex, "addr": addr, "hobby": hobby})
	})

	// 开启服务
	router.Run()
}
