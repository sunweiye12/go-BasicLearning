package main

import (
	"github.com/gin-gonic/gin"
)

/*
	如果要将收到的checkbox数据绑定到模型中，使用gin框架提供的方法
*/

type CheckBox struct {
	// 注意如果是checkbox，那么标签中的key要加[]，还有就是属性首字母一定要大写（可见性）
	Hobby []string `form:"hobby[]" json:"hobby[]" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/checkbox", func(c *gin.Context) {
		// 声明一个结构体checkbox
		var checkbox CheckBox
		// c.ShouldBind(&checkbox)通过 request 提交的 from 表单中我们可以将checkbox于对应的
		// checkbox一应上,因为在CheckBox结构体中定义了其在 form 中的 name --> hobby[]
		if c.ShouldBind(&checkbox) == nil {
			c.JSON(200, gin.H{"hobby": checkbox.Hobby})
		} else {
			c.JSON(400, gin.H{"message": "invalid request params"})
		}
	})
	router.Run()
}
