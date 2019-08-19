package main

import (
	"github.com/gin-gonic/gin"
)

/*
如果要接受多个hobby，就必须使用gin.Context.PostFormArray("key")来接收即可
*/

func main1() {
	router := gin.Default()

	router.POST("/checkbox", func(c *gin.Context) {
		// 通过PostFormArray来接收 post 传递的checkbox数据
		data := c.PostFormArray("hobby[]") //注意接收的key是一个数组后面加[]
		c.JSON(200, gin.H{"message": data})
	})

	router.Run()
}
