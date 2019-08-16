package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
可以设置响应不同的数据格式,目前支持的数据格式有JSON,XML,YAML,String,HTML
	//返回json格式
func (c *Context) JSON(code int, obj interface{})
	//返回xml格式
func (c *Context) XML(code int, obj interface{})
	//返回yaml格式
func (c *Context) YAML(code int, obj interface{})
	//返回string格式
func (c *Context) String(code int, format string, values ...interface{})
	//渲染html模板后返回
func (c *Context) HTML(code int, name string, obj interface{})

返回码,可以在返回的美容中添加返回码,一般正常情况手动指定为 200  --> 当然也可以指定 http 包下的语义化常量
*/

func main() {
	router := gin.Default()
	// 返回json
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "return json data"})
	})
	// 返回 string
	router.GET("/string", func(c *gin.Context) {
		c.String(200, "message: %s", "return string data")
	})
	// 返回 yaml 文件
	router.GET("/yaml", func(c *gin.Context) {
		arr := [][]string{ // 声明一个二维切片类型
			{"one", "two", "three"},
			{"four", "five", "six"},
		}
		c.YAML(200, arr)
	})
	// 返回 xml 文件
	router.GET("/xml", func(c *gin.Context) {
		person := struct { //声明一个匿名结构体
			Name string
			Age  int
		}{"Jane", 20}
		c.XML(200, fmt.Sprintln(person))
	})
	// 开启服务
	router.Run()
}
