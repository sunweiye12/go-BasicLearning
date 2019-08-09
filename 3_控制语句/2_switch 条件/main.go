package main

import "fmt"

/*
// switch 语句格式
// 如果有 optionalStatement 则后面的 ‘;’ 是必须的，不论 optionalValue 是否存在
// case 下的 block 语句执行完后不会再往下个 case 走所以不用末尾加 break，(***与 java 的不同点***)
// 如果要往下走则使用 fallthrough 语句

switch optionalStatement; optionalValue {
// optionalStatement：可选的声明语句，
// optionalValue ：可选得待比较值，如果未填写则默认为 true
case expression1: // 表达式 expression1 为真执行 block1
    block1
...

case expressionN:
    blockN
default: // default 可选的，上面都未匹配则执行 default 下的 blockD
    blockD
}
*/
func main() {

	// 不添加表达式默认为 true --> 直接去判断 case 条件
	switch {
	case 10 > 0:
		fmt.Println("1 条件正确,输出了(后面不执行)")
	case 1 < 0:
		fmt.Println("不输出")
	case 1 > 0:
		fmt.Println("哈哈")
	}

	// switch 带声明语句,不带表达式(**声明语句后面必须要加 ; 分号**)
	switch v := 1; {
	case v == 1:
		fmt.Println("2 条件正确,输出了")
	case v == 2:
		fmt.Println("条件不正确,不进行输出")
	default:
		fmt.Println("默认执行路径,已有符合条目,不执行")
	}

	//带有表达式(后面的 case和表达式进行匹配)
	switch v := 1; v {
	case 0:
		fmt.Println("条件不正确,不进行输出")
	case 1:
		fmt.Println("3 条件正确,输出了")
	case 2:
		fmt.Println("条件不正确,不进行输出")
	}

	// 通过 fallthrough 关键字来继续判断后面的 case
	// 语义 遇到第一个符合条件的 case,后面的不再判断,只要后面有 fallthrough 关键字就向下执行
	switch v := 1; v {
	case 0:
		fmt.Println("条件不正确,不执行")
		fallthrough
	case 1:
		fmt.Println("4 条件正确,进行输出,并向下找")
		//v++
		fallthrough
	case 2:
		fmt.Println("5 条件也正确,输出呗,并向下找")
		fallthrough
	case 3:
		fmt.Println("条件不正确,不输出1111")
		fallthrough
	default:
		fmt.Println("不输出了111")
	}

}
