package main

import (
	"fmt"
)

const c = "C" // 定义常量

var v int = 5 // 定义全局变量

//const (		//常量的声明方式,在局部变量中不能这样声明 但是可以简化版
//	haha = 12
//	hehe = 13
//)

type T struct{} //定义一个结构体类型

/*会首先执行初始化函数(在 main 之前)*/
func init() {
	fmt.Println("包结构的初始化")
}

/*程序入口*/
func main() {
	var a int = 11 //声明局部变量,如果不赋值的话默认为 0
	fmt.Println(a)
	fmt.Println(haha)
	fmt.Println(hehe)
	fmt.Println(a)
	Func1()     // 调用方法 Func1
	t := new(T) // 创建一个 T 类型对象
	t.Method1() // 调用 t 中的一个方法
}

/*构建一个方法*/
func Func1() {
	fmt.Println("函数Func1")
}

/*T结构体 来实现Method1方法,可以在 T类型的对象中调用*/
func (t T) Method1() {
	fmt.Println("函数Method1")
}
