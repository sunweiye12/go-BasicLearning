package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

/*
	输出日志文件
*/
func main() {
	//取消控制台中日志的字体颜色
	gin.DisableConsoleColor()

	//创建一个日志文件
	access_log, _ := os.Create("access_log.log")   // 在当前文件夹下面创建 access_log.log 的文件
	gin.DefaultWriter = io.MultiWriter(access_log) // 将日志信息写入到文件中

	router := gin.Default()
	router.GET("/log", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "你好数据"})
	})

	router.Run()
}
