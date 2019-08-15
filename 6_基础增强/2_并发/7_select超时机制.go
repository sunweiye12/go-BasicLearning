package main

import (
	"fmt"
	"time"
)

/*
Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。

*/

func main7() {
	c := make(chan bool)

	// 通过 select 从通道中取出元素的时候,可以设置一个超时时间,当达到超时时间的时候就会走自定义的逻辑
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second): // 当等待三秒钟还没有获取到c 中的元素时候就会执行
		fmt.Println("超时了")
	}
}
