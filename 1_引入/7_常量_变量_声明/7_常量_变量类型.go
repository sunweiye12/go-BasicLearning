package main

import (
	"fmt"
	"os"
	"runtime"
)

//程序入口函数
func main() {

	fmt.Println(Sunday)
	fmt.Println(Tuesday)
	fmt.Println(Thursday)
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a3)

	// 声明局部变量并且打印
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	// 使用 := 赋值操作符,用于变量的初始化(省去 var 关键字)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	// %s 代表字符串标识符、%v 代表使用类型的默认输出格式

}

/*常量的定义(一般用大写,如果只想在本包中使用的话可以使用_MAX 或者 cMAX 来声明)*/
/*声明格式:const identifier [type] = value
类型 type 可以省略,因为编译器会自动推导
声明的常量必须在编译器就确定
重要: 在定义常量组的时候如果不提供初始值,会默认使用上行的表达式
*/
const beef, two, c = "eat", 2, "veg"

const Monday1, Tuesday1, Wednesday1, Thursday1, Friday1, Saturday1 = 1, 2, 3, 4, 5, 6

const ( // 常量组的定义(不提供初始值时默认使用上行的表达式式) 因此每一行个数要相同,并且第一行一定要赋值哦
	Monday2, Tuesday2, Wednesday2 = 1, 2, 3
	Thursday2, Friday2, Saturday2
)

const ( // 利用 iota 关键字--> 从const 关键字开始,每定义一个常量iota就递增加1(从0开始),当遇到第一个iota赋值的常量时此值获取 iota 所对应的的值,后面的常量默认都递增加一,遇到 const重置为 0(实现一种枚举现象)
	// 相当于 iota 就代表这个 coust 常量组中的第几个常量 (从 0 开始)
	Sunday = 'a'
	Monday
	Tuesday = iota
	Wednesday
	Thursday
	Friday
	Saturday
)

/*变量的定义*/
/*声明格式:var identifier type
	Go 和许多编程语言不同，它在声明变量时将变量的类型放在变量的名称之变量被声明之后，
	系统自动赋予它该类型的零值：
	int 为 0，
	float 为 0.0，
	bool 为 false，
	string 为空字符串，
	指针为 nil。
记住，所有的内存在 Go 中都是经过初始化的。

*/

var a, b *int // 声明为指针类型的 a 和 b

// 分别声明三个类型的变量
var a1 int
var b1 bool
var str1 string

//上面的简化版(这种一般用于声明全局变量)
var (
	a2   int
	b2   bool
	str2 string
)

//声明变量并进行初始化
var (
	a3              = 15
	b3              = false
	str3            = "Go says hello to the world!"
	numShips        = 50
	city     string = "beijing"
)
