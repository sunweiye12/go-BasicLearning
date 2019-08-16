package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main1() {
	router := gin.Default()

	//	1. http://localhost:8080/jane/20/beijing/female
	/*
		// 在 relativePath 中定义了四个变量 --> http://localhost:8080/jane/20/beijing/female(上面设置了 4 个参数,所以在访问的时候必须添加四个参数)
		router.GET("/:name/:age/:addr/:sex", func(c *gin.Context) {
			// 获取 url 中传递的参数的方式
			name := c.Param("name")		// 方式 1
			fmt.Println(name)
			age := c.Params.ByName("age")	// 方式 2
			fmt.Println(age)
			// 以 json 的方式将url 传入的参数返回给客户端
			c.JSON(200, fmt.Sprintln(c.Params))
		})
	*/

	// 2. http://localhost:8080/user/jane/20/beijing/female?id=999&height=170&weight=100 (获取参数id、height、weight的值)

	router.GET("/user/:name/:age/:addr/:sex", func(c *gin.Context) {
		// 获取传递进来的参数(当查询的参数不存在时返回空的字符串,不会报错)
		fmt.Println(c.Params)
		id := c.Query("id")
		height := c.Query("height")
		weight := c.Query("weight")               // 当没有传递此参数的时候,会返回空
		city := c.DefaultQuery("city", "beijing") //当没有传入此参数时,返回默认的字符串
		c.JSON(200, gin.H{"id": id, "height": height, "weight": weight, "city": city})
	})

	// 开启服务
	router.Run()
}
