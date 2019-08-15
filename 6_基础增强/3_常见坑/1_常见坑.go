package main

import (
	"fmt"
	"time"
)

/*
1。常用的包结构
2。再利用slice切片数据类型的时候，由于他直接可以进行值传递，但是他有可以支持自动扩容，自动扩容以后地址值会发生变化
	因此当函数的参数传递为切片类型的时候，最好将操作后的切片通过return返回（因为你一旦发生了自动扩容之前的slice索引就不能用了）
*/

func main() {
	t := time.Now() // 获取当前时间
	fmt.Println(t)
	fmt.Println(t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")) // 尽量直接使用官方提供的常量,或者下面的定义常量,不要自己随便修改
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.RFC3339))

}
