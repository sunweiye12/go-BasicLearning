package visibility

import "fmt"

const PI = 3.145	// 共有常量
const pi = 3.14		// 私有常量
const _PI = 3.14	// 私有常量

var P int = 1	// 共有变量
var p int = 1	// 私有变量

/*私有函数*/
func private_Function() {
	fmt.Println("only used in this package!")
}

/*共有函数*/
func Public_fuction() {
	fmt.Println("used in anywhere!")
}
