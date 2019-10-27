package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")

	// ======调用函数并输出结果========
	a := 2
	b := 3
	fmt.Println(a, " + ", b, " = ", sum(a, b))

	// ----------开启一个线程-------------
	go test_goroute(100, 200)
	// ----------开启多个线程-------------
	for i := 0; i < 10; i++ {
		go myPrint(i)
	}
	// =========使得主线程在此此处暂停1秒=====
	time.Sleep(1 * time.Second)
}

func sum(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func test_goroute(a int, b int) {
	var sum int
	sum = a + b
	fmt.Println(sum)
}

func myPrint(a int) {
	fmt.Println(a)
}
