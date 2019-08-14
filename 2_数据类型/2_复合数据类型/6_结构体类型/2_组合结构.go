package main

import "fmt"

/*
 通过组合结构实现代码的复用
	---> 当嵌入结构中的子字段和外部结构中的字段有重复的时候的处理逻辑
		默认情况下外层的字段会覆盖内层的字段,如果想得到内层的字段a.B.Name
*/

func main2() {
	// 创建一个学生类
	a := student{
		human: human{0}, // 通过这种方式来声明切入结构的字段
		Name:  "学生",
		Age:   18,
	}
	//创建一个老师类
	b := teacher{
		human: human{1}, // 作为嵌入结构,其内部的字段已经在外部可以直接调用--> 也可以直接获取
		Name:  "老师",
		Age:   28,
	}
	b.Age = 12
	b.Sex = 100 // 可以直接获取到内嵌的字段进行复制
	fmt.Println(a, b)

	// 内层重名的逻辑 --> 内层和外层有重名的字段
	c := A{
		B:    B{"张山"},
		C:    C{"李四"},
		Name: "小 A",
	}
	// 直接 c.name 获取的是外层的字段
	fmt.Println(c.Name, c.B.Name, c.C.Name)
	// 如果外层内有此字段,内层仅有一个的字段调用时,直接获取内层字段
	// 如果外层没有此字段,而内层的多个结构中存在相同的字段,在调动是会报错
}

// 一个公共的人结构
type human struct {
	Sex int
}

// 一个学生结构
type student struct {
	human // 嵌入结构--> 组合结构
	Name  string
	Age   int
}

// 一个老师结构
type teacher struct {
	human
	Name string
	Age  int
}

// 重名字段的处理逻辑
type A struct {
	B
	C
	Name string
}

type B struct {
	Name string
}

type C struct {
	Name string
}
