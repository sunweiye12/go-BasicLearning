package main

import (
	"fmt"
)

/*
Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
*/

func main6() {

	c1 := make(chan int)

	// 开启一个线程执行匿名函数
	go func() {
		for v := range c1 {
			fmt.Println(v)
		}
	}()

	// 随机的向通道 c1 中放入 0 或者 1   (**重要在 select 只可以向通道中添加数据或者取出数据--> 如果是一个空的 select 函数会进行阻塞**)
	for i := 0; i < 5; i++ {
		select { // 随机执行
		case c1 <- 1:
		case c1 <- 0:
		}
	}
}
