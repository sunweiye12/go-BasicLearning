package main

import "fmt"

/*
指针其实就是指向一个对象（任何一种类型数据、包括指针本身）的地址值，
对指针的操作都会映射到指针所指的对象上
*/

func main() {

	var p *int // 定义指向int型的指针，默认值为空：nil
	// nil指针不指向任何有效存储地址，操作系统默认不能访问
	// fmt.Printf("%x\n", *p) // 编译报错

	var a int = 10
	p = &a        // 取地址(指针 p 取到整型 a 的地址值)
	add := a + *p // 声明变量 add(对指针的操作会映射到对象身上)

	fmt.Println(a)   // 输出：10
	fmt.Println(p)   // 输出：0xc0420080b8
	fmt.Println(add) // 输出：20
}