package main

import (
	"fmt"
	"time"
)
/*
说到通道 channel，则必须先了解下 Go 语言的 goroutine 协程（轻量级线程)。
channel 就是为 goroutine 间通信提供通道。
goroutine 是 Go 语言提供的语言级的协程，是对 CPU 线程和调度器的一套封装。
channel 也是类型相关的，一个 channel 只能传递一种类型的值。(可以用于多个线程之间的通信)
声明：var 通道名 chan 通道传递值类型
或
	make 函数初始化：make(chan 值类型, 初始存储空间大小)
*/

func main() {
	var ch1 chan int            // 声明一个通道
	ch1 = make(chan int)        // 未初始化的通道不能存储数据，初始化一个通道
	ch2 := make(chan string, 2) // 声明并初始化一个带缓冲空间的通道

	// 通过匿名函数向通道中写入数据，通过 <- 方式写入
	go func() { ch1 <- 1 }()
	go func() { ch2 <- `a` }()
	go func() { ch2 <- `b` }()

	v1 := <- ch1 // 从通道中读取数据
	v2 := <- ch2	// 类似一个栈结构:先入后出
	v3 := <- ch2
	fmt.Println(v1) // 输出：1
	fmt.Println(v2) // 输出：b
	fmt.Println(v3) // 输出：a

	// 写入，读取通道数据
	ch3 := make(chan int, 1) // 初始化一个带缓冲空间的通道

	//通过 go 来开启两个协程,使得这两个方法并发执行
	go readFromChannel(ch3)
	go writeToChannel(ch3)

	// 开启的这两个协程相当于是主线程的子线程,相对独立,但必须保障主线程没有结束,他们才会执行
	// 如果主线程结束了,协程也会结束,因此下面才会休眠
	// 主线程休眠1秒，让出执行权限给子 Go 程，即通过 go 开启的 goroutine，不然主程序会直接结束
	time.Sleep(1 * time.Second)		// 调用time.Sleep()来时主线程陷入沉睡 time.Second指1 秒
}

func writeToChannel(ch chan int) {
	for i := 1; i < 10; i++ {
		fmt.Println("写入：", i)
		ch <- i
	}
}

func readFromChannel(ch chan int) {
	for i := 1; i < 10; i++ {
		v := <- ch
		fmt.Println("读取：", v)
	}
}