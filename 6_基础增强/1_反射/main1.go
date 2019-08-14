package main

import (
	"fmt"
	"reflect"
)

/*
反射会将匿名字段作为独立字段（匿名字段本质）
想要利用反射修改对象状态，前提是 interface.data 是 settable，
即 pointer-interface
- 通过反射可以“动态”调用方法
*/

func main() {
	// 实例化一个结构
	u := User{
		Id:   1,
		Name: "张山",
		Age:  10,
	}

	fmt.Println(u)
	//通过反射修改类型字段(传入的是地址值)
	Set(&u)
	fmt.Println(u)

	// 正常调用方法
	u.Hello("李四")
	// 通过反射来调用方法
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("joe")} // 声明一个参数
	mv.Call(args)

}

// 声明一个结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 传入一个结构体,通过反射修改内容
func Set(o interface{}) {

	v := reflect.ValueOf(o)                            // 获取类型的详细信息
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { // v.Kind()==reflect.Ptr 判断 v 是否为指针类型 && 并且不能被修改,则返回错误
		fmt.Println("有错误")
		return
	} else {
		v = v.Elem() // 如果是指针类型获取此类型所指向的字段
	}

	f := v.FieldByName("Name") //获取一个名字为 Name 的字段
	if !f.IsValid() {          // 如果此字段不存在的话 直接返回错误
		fmt.Println("字段不存在")
		return
	}

	if f.Kind() == reflect.String { //如果存在这个字段,并且字段的类型为 string
		f.SetString("NBA")
	}
}

// 为 User 结构体声明一个方法
func (u User) Hello(name string) {
	fmt.Println("hello", name, ", my name is", u.Name)
}
