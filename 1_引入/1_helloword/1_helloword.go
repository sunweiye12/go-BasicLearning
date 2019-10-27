package main

/*
必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。
package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
*/
/*go语言的基本结构和要素*/
//单行注释
import (
	"fmt"
)

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("这是第一个go程序")
	test_pipe()
}

// 演示channel的执行流程
func test_pipe() {
	// 创建并添加数据
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	// pipe  <- 4  // 如果沟道中的数据已经加满了再往里面添加数据就会导致死锁(导致线程阻塞)
	fmt.Println(len(pipe)) // 打印沟道的长度

	// 去除数据并打印
	var t1 int
	t1 = <-pipe // 在管道中取出来一个值赋值给t1
	fmt.Println(t1)
}
