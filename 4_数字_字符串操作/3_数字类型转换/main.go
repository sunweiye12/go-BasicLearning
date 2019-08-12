package main

import (
	"fmt"
	"strconv"
)

/*
	整型从高位类型转低位类型会有精度丢失
	浮点型转整型会丢失小数点后的值
	复数型转非复数整型时符号丢失。
		数据类型转换格式：目标类型(转换类型)

	使用 strconv 包中定义的函数做数字和字符串转换。
*/
func main() {
	// 整型转浮点型
	var i int = 1
	fmt.Printf("%f\n", float32(i)) // 输出：1.000000

	// 浮点型转整型
	var f float32 = 3.1415926
	fmt.Printf("%d\n", int(f)) // 输出：3，小数后丢失

	// float32 转 float64
	fmt.Printf("%v\n", float64(f)) // 输出：3.141592502593994，6位后的小数精度是错误的

	// float64 转 float32
	var f2 float64 = 3.141592653589793
	fmt.Println("%v\n", float32(f2)) // 输出：3.1415927，精度降低

	// 将 int 转成 string
	var i1 int = 111
	s := strconv.Itoa(i1) // 数字转换成字符串
	fmt.Println(s)

	// 将string 转换成 int
	i2, _ := strconv.ParseInt(s, 10, 64) // 把 s 转为10进制64位数
	fmt.Println(i2)
}
