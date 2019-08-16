package main

import (
	"github.com/gin-gonic/gin"
)

/*
	使用浏览器向服务器发送一个文件(POST)
前端表单中要设成multipart/form-data
*/
func main() {

	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file_1, _ := c.FormFile("myfile_1")
		file_2, _ := c.FormFile("myfile_2")
		file_3, _ := c.FormFile("myfile_3")
		c.IndentedJSON(200, gin.H{
			"file_1": file_1.Filename,
			"file_2": file_2.Filename,
			"file_3": file_3.Filename,
		})
	})

	// 开启服务
	router.Run()
}
