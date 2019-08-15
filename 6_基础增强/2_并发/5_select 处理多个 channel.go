package main

import (
	"fmt"
)

/*
Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
通过利用多核的 cpu 实现并发的访问
*/

func main() {

	c1, c2 := make(chan int), make(chan string) // 创建两个 channel 用于多线程通信

	mk := make(chan bool, 2) //创建一个标志位(缓存为2)

	// 开启一个线程执行匿名函数
	go func() { // 函数的功能是不断地从 c1 和 c2 中获取数据直到都为空
		a, b := false, false
		for { // 一个无线循环,不断的在 channel 中获取数据
			select {
			case v, ok := <-c1: // 当 c1 关闭以后此处还是赋值操作,得到的 ok 为空-->就会执行 else 中的语句,因此应该加个标志,当通道为空的时候竟不再望 wk'中添加元素
				if ok { //如果 c1 中存在元素,则打印元素
					fmt.Println("c1: ", v)
				} else {
					if !a {
						mk <- true
						a = true
					}
					break
				}

			case v, ok := <-c2:
				if ok { //如果 c2 中存在元素,则打印元素
					fmt.Println("c2: ", v)
				} else {
					if !b {
						mk <- true
						b = true
					}
					break
				}
			}
		}
	}()

	c1 <- 1
	c1 <- 2
	c1 <- 3
	c2 <- "h"
	c2 <- "e"
	c2 <- "l"
	c2 <- "l"
	c2 <- "o"

	// 一定要关闭通道,否则在沟道中不断取出数据的线程会一直等待
	close(c1)
	close(c2)

	// 只有两个元素都取出来的时候(证明两个对列 c1 和 c2 都为空)
	for i := 0; i < 2; i++ {
		<-mk
	}
}
