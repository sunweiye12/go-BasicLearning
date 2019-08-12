package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	callback(1, Add)
	// 函数的返回值是讲一个函数,此函数指向 p2
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// 次函数的返回值同样也是一个函数
	where()
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

}

// 此处参数为一个函数(在调用的时候应该注意)
func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

// 所谓闭包就是类似于匿名函数的声明和调用 func(x, y int) int { return x + y }  这样写是错的
//但可以被赋值于某个变量，即保存函数的地址到变量中：
// fplus := func(x, y int) int { return x + y }，然后通过变量名对函数进行调用：fplus(3,4)。
// 或者 : 直接对匿名函数进行调用：func(x, y int) int { return x + y } (3, 4)
// 将函数作为返回值进行操作
func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

// 通过过使用闭包来实现一个来封装一个调试的函数
where := func () {
	_, file, line, _ := runtime.Caller(1)	// 用于获取当前程序运行的状态
	log.Printf("%s:%d", file, line)
}
