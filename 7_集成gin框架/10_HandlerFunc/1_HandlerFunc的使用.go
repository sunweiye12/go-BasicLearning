package main

import (
	"github.com/gin-gonic/gin"
)

/*
绑定的func(c *gin.Context){ code } 就是gin.HandlerFunc，这里是一个匿名函数的形式，
其实我们完全可以将它提出来，让多个路径共用，比如下面这样
*/

//声明一个gin.HandlerFunc
func response(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello world"})
}

func main() {
	router := gin.Default()
	// 当访问下面热加一路径的时候都会调用上面定义的 response 函数
	router.GET("/demo", response)
	router.GET("/test", response)
	router.GET("/aaaa", response)
	router.Run()
}
