package main

import "fmt"

/*
select 是用于通道 channel 通信的选择语句，格式：
// select 语句，用于通道通信选择
select {
case sendOrReceive1: // 发送或接受通道信息
    block
...

case sendOrReceiveN:
default:
    blockD
}
*/
func main() {
	// 声明并初始化三个通道类型
	// make() 函数用于声明slice切片、map字典、channel通道。
	ch1 := make(chan int)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 2)

	//ch1 <- 2
	ch2 <- 2 // 往通道 ch2 中写入一个数据

	select {
	case i1 := <-ch1:
		fmt.Println("读取 ch1 数据：", i1)
	case i2 := <-ch2:
		fmt.Println("读取 ch2 数据：", i2) // 输出：读取 ch2 数据： 2
	case ch3 <- 3:
		fmt.Println("写入 ch3 数据：3")
	default:
		fmt.Println("nothing")
	}
}
