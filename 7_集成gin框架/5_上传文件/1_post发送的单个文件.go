package main

import (
	"github.com/gin-gonic/gin"
)

/*
	使用浏览器向服务器发送一个文件(POST)
前端表单中要设成multipart/form-data
*/
func main1() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("myfile") //获取文件
		filename := file.Filename       // 获取文件名
		size := file.Size               // 获取文件大小
		header := file.Header           // 获取文件头信息
		c.JSON(200, gin.H{
			"filename": filename,
			"size":     size,
			"header":   header,
		})
	})

	// 开启服务
	router.Run()
}
