package main

import "fmt"

/*
 方法的概念和函数类似,值不过方法值隶属于一个结构体
因为在声明一个结构体的时候只包含可字段,而没有结构体的方法,---> 所以在 go 语言中结构体的方法声明是在外部通过方法来实现的
----> 函数可以直接使用,而方法是属于结构体,只能够通过结构体的实例来哦实现调用
*/

func main() {

	a := A{Name: "占山"}
	a.Print() // 调用 A 中所对应的的方法

	b := B{Name: "李四"}
	b.Print() // 调用 B 中所对应大方法
	//A.Print()

	b.Print1() // 调用 B指针参数的方法,在方法内改变对象将直接引起原对象的变化(不需要用地址值调用,可以直接对象来调用)
	fmt.Println(b)

	// 自定义类型
	var c ZE = 1
	// 调用自定义类型中增强的方法
	c.Incr(100)
	fmt.Println(c)

}

// 两个对象 A 和 B
type A struct {
	Name string
}

type B struct {
	Name string
}

//虽然在 go 语言中不支持方法的重载.但是下面这两个方法虽然方法名相同,当时却不属于同一个结构,因此不在方法重载的范涛
//声明一个 A 结构的方法
func (a A) Print() {
	fmt.Println(a)
}

//声明一个 B 结构的方法
func (b B) Print() {
	fmt.Println(b)
}

// 在声明方法的时候也尽量引入指针对象---> 这样就可以直接基于对象来进行修改(影响到传入的原对象)
func (b *B) Print1() {
	b.Name = "王五"
	fmt.Println(b)
}

// 声明一个类型为 int 的结构类型(可以认为是类型别名,区别点在于可以为ZE 类型定义专属方法)
type ZE int

// 为基本数据类型定义属于自己的方法(每次调用此方法,都会自增num 的值)---> 传入的是地址值
func (c *ZE) Incr(num int) {
	*c += ZE(num) //c 里面传入的是地址值,通过*c 得到此地址上的值,因而ZE 和 int 不属于同一种数据类型,因此在运算的时候将 num 转换成 ZE 雷静进行运算
}
