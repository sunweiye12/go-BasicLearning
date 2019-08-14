package main

import "fmt"

/*
约定俗成
	1.main() 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）。
	2.好的程序是非常注意DRY原则的，即不要重复你自己（Don't Repeat Yourself）
	3.return 语句也可以用来结束 for 死循环，或者结束一个协程（goroutine）。
	4.格式要求 func g() {  }  左面大括号一定要与函数名在同一行 -- 否则报错
	5.在go 语言中不支持方法的重载
	6.假设 f1 需要 3 个参数 f1(a, b, c int)，同时 f2 返回 3 个参数 f2(a, b int) (int, int, int)
		就可以这样调用 f1：f1(f2(a, b))。
	7.返回值赋给了空白符 _，默认自动丢弃
	8.关键字 defer(修饰某个语句或函数) 的用法类似于Java 的finally 语句块，它一般用于释放资源
	9.不支持重载,嵌套
*/
var num int = 10
var numx2, numx3 int

func main() {

	// 调用多个返回值的函数
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()

	// 将第二个返回值抛弃, 只要一个
	numx2, _ = getX2AndX3_2(num)
	PrintValues()

	// 变长参数函数
	bianchang(1, 2, "3", "4", "5")

	// 测试 defer 关键字
	function1()
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

// 多个返回值的函数
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 一般使用上面的用法
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

/*
函数中的变长参数
函数的最后一个参数是采用 ...type 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0
由接片类型实现(此处注意,因为切片类型属于引用类型,如果在函数体改变会影响到原数据)
*/
func bianchang(a int, b int, c ...string) {
	fmt.Println("这是一个变长参数")
	fmt.Print(a, " ")
	fmt.Print(b, " ")
	// 遍历元组进行打印
	for _, v := range c {
		fmt.Print(v, " ")
	}
}

// defer 关键字的使用 defer(修饰某个语句或函数) 的用法类似于Java 的finally 语句块，它一般用于释放资源
func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2() // 在本方法结束以前调用 defer 修饰的方法
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}
