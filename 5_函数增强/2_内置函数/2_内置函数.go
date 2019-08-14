package main

import (
	"fmt"
)

/*
	Go 中没有 class 而引入了 struct结构它与 class 很相似,但在 go 中不存在继承的概念。

*/

func main() {

	// 创建一个person结构实例 --> 此时得到的是对象的值
	a := person{
		Name: "zhangsan",
		Age:  0,
	}
	fmt.Println(a)

	// 调用函数
	A(a)   // 证明此处是值传递(将原对象进行了复制进行传递,没有改变原值)
	_A(&a) // 要想实现地址值传递,可以可通过将地址值传入达到此效果  (** 一般处于性能的考虑都会进行地址值的传递 **)
	fmt.Println(a)

	// 此处得到的是对象地址值(直接获取的是指针)  -- >   (**推荐创建结构的时候直接创建地址值 **)
	b := &person{
		Name: "lisi",
		Age:  0,
	}
	fmt.Println(b)
	_A(b)
	fmt.Println(b)

	// 构建一个匿名结构 --> 临时定义的一个结构别的地方不会用到,所以采用匿名方式 --> 两个大括号
	c := &struct {
		Name string
		Age  int
	}{
		Name: "wangwu",
		Age:  99,
	}
	fmt.Println(c)

	// 匿名字段
	type student struct { // 字段没有声明自己的名字
		string
		int
	}
	d := student{ //在声明对象时必须奥严格与匿名字段的类型对应上
		string: "张山",
		int:    88,
	}
	fmt.Println(d)
}

// 定义了一个空的结构体
type test struct {
}

// 定义一个 person 的结构体
type person struct {
	Name string
	Age  int
}

// 创建一个函数 A  ---> 传入一个 person 对象将 Age 属性变为 13,然后将此对象打印
func A(per person) {
	per.Age = 13
	fmt.Println("A", per)
}

func _A(per *person) {
	per.Age = 13
	fmt.Println("A", per)
}
