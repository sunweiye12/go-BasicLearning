package main

import (
	"fmt"
	"runtime"
	"sync"
	//"runtime"
)

/*
Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
通过电泳多核的 cpu 实现并发的访问
*/

func main4() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // 调用电脑 cpu 核数的 cpu 来执行任务(调用多核 cpu 来处理资源)

	// --------通过通道缓存来实现同步并发功能
	c := make(chan bool, 5)              // 开创一个缓存为 5 的 channel 类型
	for index := 0; index < 5; index++ { //开辟 5条线程来执行 Go 函数(每个函数向channel里面添加一个元素)
		go Go(c, index)
	}

	// 在 channel 中取出 5个元素后(说明 5 个Go函数都已经执行完了)然后结束主线程
	for i := 0; i < 5; i++ {
		<-c
	}
	fmt.Println("----------------分界线----------------------")

	// -----通过同步包来实现同步并发功能
	wg := sync.WaitGroup{} // 创建一个同步包
	wg.Add(5)              // 为同步包添加 5 个任务

	for i := 0; i < 5; i++ { // 开启 5 个线程来并发执行
		go Go2(&wg, i) // 传入包的引用地址值
	}

	wg.Wait() // 等待上面所有任务执行完成(wg中设置的任务都完成以后再向下执行)
}

// 此函数每次执行完成以后会在 channel 中放入一个元素
func Go(c chan bool, index int) {

	a := 1
	for i := 0; i < 1000000; i++ {
		a = +i
	}
	fmt.Println(index, a)
	c <- true // 执行完成以后将一个传入 channel
}

//  每次执行完成都减少一个任务量
func Go2(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a = +i
	}
	fmt.Println(index, a)

	wg.Done() // 减少一个任务量
}
