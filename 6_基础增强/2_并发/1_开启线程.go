package main

import (
	"fmt"
	"time"
)

/*
Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
Channel
	Channel 是 goroutine 沟通的桥梁，大都是阻塞同步的
	通过 make 创建，close 关闭
	Channel 是引用类型
	可以使用 for range 来迭代不断操作 channel
	可以设置单向或双向通道
	可以设置缓存大小，在未被填满前不会发生阻塞

Select
	可处理一个或多个 channel 的发送与接收
	同时有多个可用的 channel时按随机顺序处理
	可用空的 select 来阻塞 main 函数
	可设置超时
*/

func main1() {

	// 开启一个 goroutine 线程,来执行一个函数
	go Go1()

	// 使主线程沉睡 2 秒,使得基于主线程开启的 goroutine 能够顺利执行完成
	time.Sleep(1 * time.Second)
}

func Go1() {
	fmt.Println("Go Go Go !!!")
}
