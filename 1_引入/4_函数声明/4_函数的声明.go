package main

import (
	"fmt"
	"math"
)

/*定义文件的入口*/
func main() {
	var a int = 11 //如果不赋值的话默认为 0
	fmt.Println(a)
	// 调用无参函数
	fristFunc()
	// 调用有参函数
	b := sum(3, 4)
	fmt.Println(b)

	// 调用多值返回函数,用于判断函数是否成功(如果输入参数为小于0的时候err返回为空,只有大于0时才为true )
	c, err := mySqrt(4)
	if err == true { //当 err 为 true 时表明函数正确执行
		//fmt.Println(err)
		fmt.Println(c)
	}
	// 调用多值返回函数,由于第二个参数没有用,这里定义为_(空),会被直接丢弃掉
	i1, _, f1 := ThreeValues()
	fmt.Printf("The int: %d, the float: %f \n", i1, f1)
}

/*第一个基本的函数(私有)*/
func fristFunc() {
	//println("abc")  也可以打印不过一般用于阶段性测试
	fmt.Println("第一个函数")
	fmt.Println("第二个有参数的函数", sum(2, 5))
	//fmt.Println(sum(2,5))
}

/*第二个有参数的函数,返回类型为 int (私有)*/
func sum(a int, b int) int {
	return a + b
}

/*
	多返回值的函数调用
Go 语言的函数经常使用两个返回值来表示执行是否成功：(为我们判断一个函数是否正常执行)
	返回某个值以及 true 表示成功；
	返回零值（或 nil）和 false 表示失败
*/

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	} // 发生错误的 case
	return math.Sqrt(f), true
}

func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}
