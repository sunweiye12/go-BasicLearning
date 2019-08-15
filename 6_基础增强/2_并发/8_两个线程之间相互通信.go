package main

import "fmt"

// 声明一个用于通信的线程
var c chan string

func main() {
	c = make(chan string) // 初始化
	go Pingpong()         // 开启线程执行内部任务
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("From Main: Hi, #%d", i) // 2 向通道中添加一个数据并阻塞,之前阻塞的线程呢个就开启  (Sprintf将格式化之后的输出语句转化成字符串)
		fmt.Println(<-c)                          // 4 由于通道为空,所以进入阻塞     //6 县线程中存在数据,开启线程后取出数据
	}
}

// 声明一个函数
func Pingpong() {
	i := 0
	for {
		fmt.Println(<-c)                              // 1 在通道中取出一个数据,因为现在为空,所以此线程阻塞      3 通道中有数据,开启线程从通道中取出数据释放上面阻塞线程,此时通道又为空
		c <- fmt.Sprintf("From Pingpong: Hi, #%d", i) // 5 向通道中添加一个数据,并阻塞
		i++
	}
}
