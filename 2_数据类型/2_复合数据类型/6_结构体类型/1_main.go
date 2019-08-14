package main

import (
	"fmt"
)

/*
go 中不存在 class 结构而是存在 struct的数据类型代替 class 的功能
*/

func main1() {
	// 创建 person 对象
	a := person{
		Name: "张山",
		Age:  12,
	}
	fmt.Println(a)

	//调用函数 A  --> 由于是指传递,因此不改变原值
	A(a)
	fmt.Println(a)
	//调用函数 -A -->由于传递的是地址值,因此会改变原值
	_A(&a)
	fmt.Println(a)

	// 因此我们在声明对象时可以直接得到对象的地址值 --> 一般推荐折磨来用
	b := &person{
		Name: "李四",
		Age:  18,
	}
	fmt.Println(b)
	_A(b)
	fmt.Println(b)

	// 匿名结构 --> 只在此处声明一次的结构
	c := &struct {
		Name string
		Age  int
	}{
		"王五",
		17,
	}
	fmt.Println(c)

	// 匿名字段 --> 在声明对象的时候一定要字段赋值类型不能错
	type nam struct {
		string
		int
	}
	d := &nam{
		string: "赵柳",
		int:    12,
	}
	fmt.Println(d)

}

// 定义了一个 person 的结构体
type person struct {
	Name string
	Age  int
}

// 定义一个函数 A,将传入的 person 对象的 age 属性变为 100,并打印
func A(per person) {
	per.Age = 100
	fmt.Println(per)
}

// 传入对象的地址值,--> 一般推荐这样来做
func _A(per *person) {
	per.Age = 100
	fmt.Println(per)
}
