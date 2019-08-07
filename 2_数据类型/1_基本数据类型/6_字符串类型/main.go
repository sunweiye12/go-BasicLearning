package main

import "fmt"

/*
Go 语言默认编码都是 UTF-8
*/
func main() {
	var str1 string // 默认值为空字符串 ""
	str1 = `hello world`
	str2 := "你好世界" 	//声明并初始化
	str := str1 + " " + str2 // 字符串连接
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str) // 输出：hello world 你好世界

	// 遍历字符串
	l := len(str)
	for i := 0; i < l; i++ {
		chr := str[i]
		fmt.Println(i, ":",chr) // 输出字符对应的编码数字
	}
}