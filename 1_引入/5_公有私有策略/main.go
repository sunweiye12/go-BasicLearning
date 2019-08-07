package main
/*
Go语言没有像其它语言一样有public、protected、private等访问控制修饰符,
它是通过字母大小写来控制可见性的，
如果定义的常量、变量、类型、接口、结构、函数等
	名称是大写字母开头表示能被其它包访问或调用（相当于public），
	非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头）
*/
import (
	"fmt"
	// shorturl/model 导入包文件
	visibility "testProject/1_引入/5_公有私有策略/test"
)

func main() {
	visibility.Public_fuction() //used in anywhere!
	//visibility.private_function() //不能访问私有函数，无法通过编译
	fmt.Println(visibility.P) //1
	//fmt.Println(visibility.p) //不能访问私有变量，无法通过编译
	fmt.Println(visibility.PI) //3.14
	//fmt.Println(visibility.pi) //不能访问私有常量，无法通过编译
	//fmt.Println(visibility._PI) //不能访问私有常量，无法通过编译
}
