package main
/*
一个类型的值可以被转换成另一种类型的值。由于 Go 语言不存在隐式类型转换，
因此所有的转换都必须显式转换
*/
import (
	"fmt"
)


//程序入口函数
func main() {
	//类型B的值 = 类型B(类型A的值)
	var a = 13.5			//声明的局部变量为 float 类型
	var haha int = int(a)	//通过类型转换为 int 类型
	fmt.Println(a)
	fmt.Println(haha)
}
