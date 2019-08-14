package main

import (
	"fmt"
)

/*
Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
Channel
	Channel 是 goroutine 沟通的桥梁，大都是阻塞同步的
	通过 make 创建，close 关闭
	Channel 是引用类型(所以可以直接传递)
*/

// 不用 sleep 的解决方案
func main2() {

	// 通过创建 channl 来进行线程之间的通信
	//创建一个 channel 类型
	c := make(chan bool) // 通道中存储的类型为 bool 类型
	go func() {          // 执行匿名函数(将 true 存入沟道 c)
		fmt.Println("Go Go Go!!!")
		c <- true
	}()

	// 在这里读数据,如果没有数据就会在此处阻塞,通道 c 中有元素的时候,结束阻塞并并取出元素
	<-c
	//因为程序结束所有资源都会被释放,所以不用调用 close 方法
}
