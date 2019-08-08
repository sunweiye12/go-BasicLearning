package test

import "fmt"

const PI2 = 2  // 共有常量
const pi2 = 2  // 私有常量
const _PI2 = 2 // 私有常量

var P2 int = 2 // 共有变量
var p2 int = 2 // 私有变量

/*私有函数*/
func private_Function2() {
	fmt.Println("only used in this package2!")
}

/*共有函数*/
func Public_fuction2() {
	fmt.Println("used in anywhere2!")
}
