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

	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second): // 当等待三秒钟还没有获取到c 中的元素时候就会执行
		fmt.Println("超时了")
	}
}
