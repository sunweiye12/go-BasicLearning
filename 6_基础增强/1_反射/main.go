package main

import (
	"fmt"
	"reflect"
)

/*
反射可大大提高程序的灵活性，使得 interface{} 有更大的发挥余地
反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
*/
func main1() {
	// 实例化一个结构
	u := User{
		Id:   1,
		Name: "张山",
		Age:  18,
	}

	// 调用此方法,反射出其所有的字段和方法
	Info(u)
}

// 声明一个结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 为 User 结构定义第一个方法
func (u User) Hello1() {
	fmt.Println("Hello World1")
}

// 为 User 结构定义第二个方法
func (u User) Hell2() {
	fmt.Println("Hello World2")
}

// 定义一个反射的方法
func Info(o interface{}) {
	t := reflect.TypeOf(o) // 通过反射获取传入参数的类型
	fmt.Println(t.Name())

	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
