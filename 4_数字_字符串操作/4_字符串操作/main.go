package main

import (
	"fmt"
	"strconv" // 用于字符串和数字之间的转换
	"strings"
)

func main() {
	str_int() //将字符串转化为 int
	stringDoing()
}

func str_int() {
	var str = "111"
	i, _ := strconv.Atoi(str) // 字符串类型转换为 int 类型
	fmt.Printf("%d\n", i)     // 输出：111
}

//字符串操作的函数
func stringDoing() {
	str := " hahahaha en woh"
	str1 := strings.ToUpper(str) // 转化为大写  ( 对应的为有ToLower() )
	str2 := strings.ToTitle(str)
	str3 := strings.Trim(str, "h") // 将两端的指定字符消去
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
