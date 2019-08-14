package main

import (
	"fmt"
	"strconv" // strconv.Itoa() 用于将数字转换成字符串
)

/*
接口用来定义行为。相当于多个方法签名的集合
Go 语言不同于面向对象语言，没有类的概念，也没有传统意义上的继承。
Go 语言中的接口，用来定义一个或一组行为，某些对象实现了接口定义的行为，则称这些对象实现了（implement）该接口，类型即为该接口类型。
定义接口也是使用 type 关键字，格式为：
	// 定义一个接口
	type InterfaceName interface {
		FuncName1(paramList) returnType
		FuncName2(paramList) returnType
		...
	}
*/

// 定义一个 Person 接口
type Person interface {
	Say(s string) string // 每一个行为可以有传入的参数和返回的值
	Walk(s string) string
}

// 定义一个 Man 结构体(相当于类)
type Man struct {
	Name string
	Age  int
}

// Man 实现并重写 Say 方法
func (m Man) Say(s string) string {
	return s + ", my name is " + m.Name
}

// Man 实现并重写 Walk 方法，--> strconv.Itoa() 数字转字符串****
func (m Man) Walk(s string) string {
	return "Age: " + strconv.Itoa(m.Age) + " and " + s
}

func main1() {
	var m Man       // 声明一个类型为 Man 的变量
	m.Name = "Mike" // 赋值
	m.Age = 30
	fmt.Println(m.Say("hello"))    // 输出：hello, my name is Mike
	fmt.Println(m.Walk("go work")) // 输出：Age: 30 and go work

	jack := Man{Name: "jack", Age: 25} // 初始化一个 Man 类型数据
	fmt.Println(jack.Age)
	fmt.Println(jack.Say("hi")) // 输出：hi, my name is jack
}
