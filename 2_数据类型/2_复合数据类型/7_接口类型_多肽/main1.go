package main

import (
	"fmt"
)

/*
1.接口是一个或多个方法签名的集合
2.只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示(******)  --> 当实现了一个接口以后就可以有多肽的应用
3.接口只有方法声明，没有实现，没有数据字段
4.接口可以匿名嵌入其它接口，或嵌入到结构中
5.接口调用不会做receiver的自动转换
6.接口也可实现类似OOP中的多态
7.空接口可以作为任何类型数据的容器
*/

func main() {
	// 创建一个结构实例
	a := Phone{name: "苹果手机"}
	// 调用方法 Name 返回其方法
	strName := a.Name()
	fmt.Println(strName)
	// 调用 Connect 方法
	a.Connect()

	// 调用一个传入参数为 USB 类型的函数,我们传入 a   ---->  说明 a 实现了 USB 接口
	// 此处出入的参数为 USB参数,却传入的是 Phone 类型,类似多肽的应用
	Disconnect(a)
	// 下面这个函数可以传递任意类型的参数
	MyMethod(a)
	MyMethod(1)
	MyMethod(1.1)
}

// 定义一个接口类型
type USB interface {
	Name() string
	Connecter // 引入了一个嵌入接口,将此接口中的方法引入(可以作为此接口的方法)
}

// 嵌入接口 --> 将这个接口嵌入到 USB 接口中,这样的话经相当于 USB 接口中拥有了此接口的方法
type Connecter interface {
	Connect()
}

// 定义一个结构体  (当一个结构实现了一个借口中的所有方法是,就默认实现了此接口)
type Phone struct {
	name string
}

// phone结构实现此方法(并且自己定义方法体)
func (pc Phone) Name() string {
	return pc.name
}

func (pc Phone) Connect() {
	fmt.Println(pc.name)
}

// 定义一个函数传入的参数为 USB 接口类型
func Disconnect(usb USB) {
	// 进行类型判断 (判断 usb 是否属于 Phone 类型)
	pc, ok := usb.(Phone) // 判断传入的 usb 是否为 Phone 类型
	if ok {
		fmt.Println(pc.name)
	} else {
		fmt.Println("未知设备")
	}
}

// 定义一个函数,传入的参数是任意类型,可以将参数类型这是为空接口,因为所有结构体都默认实现了空接口
func MyMethod(usb interface{}) {
	// 进行类型判断的高效方式
	switch v := usb.(type) {
	case Phone:
		fmt.Println("这是一个手机类型", v.name)
	case int:
		fmt.Println("这是一个数值类型")
	default:
		fmt.Println("这既不是 Phone 类型,也不是 int 类型")
	}
}
