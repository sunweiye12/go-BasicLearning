package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
其实在网站访问过程中，有很多的资源都是静态的，也就是说，这些静态资源是可以直接从硬盘上读取之后返回用户，
不需要服务器端在做多余的处理
*/

// gin框架中提供了三个方法来实现：
func main() {
	router := gin.Default()

	//设置静态文件目录，如果访问localhost:8080/assets/test.txt
	//如果./assets/test.txt文件存在，那么就返回该文件，否则返回404
	router.Static("/assets", "./assets") // 使用相对路径

	//和上面的功能一样
	router.StaticFS("/my_file", http.Dir("my_file"))

	//为单个文件绑定路由
	//可以通过访问localhost:8080/test.txt 来获取./assets/test.txt文件
	router.StaticFile("/test.txt", "./assets/test.txt")

	router.Run()
}
