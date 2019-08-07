package main

import (
	"fmt"
)

/*定义文件的入口*/
func main() {
	var a int = 11  //如果不赋值的话默认为 0
	fmt.Println(a)
	fristFunc()
}

/*第一个基本的函数(私有)*/
func fristFunc(){
	//println("abc")  也可以打印不过一般用于阶段性测试
	fmt.Println("第一个函数")
	fmt.Println("第二个有参数的函数", sum(2,5))
	//fmt.Println(sum(2,5))
}

/*第二个有参数的函数,返回类型为 int (私有)*/
func sum(a int, b int) int {
	return a + b
}

