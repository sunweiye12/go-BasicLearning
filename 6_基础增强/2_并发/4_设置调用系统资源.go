package main

import (
	"fmt"
	"runtime"
	//"runtime"
)

/*
Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
通过电泳多核的 cpu 实现并发的访问
*/

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // 调用电脑 cpu 核数的 cpu 来执行任务(调用多核 cpu 来处理资源)

	c := make(chan bool, 5)
	for index := 0; index < 5; index++ {
		go Go(c, index)
	}
	for i := 0; i < 5; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {

	a := 1
	for i := 0; i < 1000000; i++ {
		a = +i
	}

	fmt.Println(index, a)

	c <- true

}
