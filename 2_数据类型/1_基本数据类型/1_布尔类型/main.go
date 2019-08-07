package main

import "fmt"

/*布尔类型 (bool)
值：true 和 false，默认值为 false*/

func main() {
	var v1, v2 bool         // 声明变量，默认值为 false
	v1 = true               // 赋值
	v3, v4 := false, true   // 声明并赋值(更简单的初始化方式)

	fmt.Println("v1:", v1)   	// v1 输出 true
	fmt.Println("v2:", v2) 	// v2 没有重新赋值，显示默认值：false
	fmt.Println("v3:", v3) 	// v3 false
	fmt.Println("v4:", v4) 	// v4 true
}