package main

import (
	"github.com/gin-gonic/gin"
)

/*
命令窗口: go get -u github.com/gin-gonic/gin 安装gin 框架
默认绑定的是 8080 端口,可以通过访问 http://localhost:8080/get 过去服务
*/
func main() {
	router := gin.Default()
	router.GET("/get", // 设置访问的方式和路径
		func(c *gin.Context) { // 设置访问返回的信息
			c.JSON(200, gin.H{"message": "你好呀世界", "这是什么": 1}) // gin.H{ }这只是一个map结构，别以为是一个struct哈 --> type H map[string]interface{} value 可以为任何类型
		})
	//router.Run()		// 应用默认绑定 8080 端口
	router.Run(":9000") // 将应用端口改为 9000
}
