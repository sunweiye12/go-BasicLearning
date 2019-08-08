package main

import "fmt"

/*
// if 语句格式
// else if 和 else 都是可选的，且 else if 可以有多个
// optionalStatement1 声明语句如（v := xx）中定义的变量其作用域只在 if ... else if ... else 范围内有效
if optionalStatement1; booleanExpression1 { // optionalStatement1：可选的声明语句，booleanExpression1：真假表达式，为真执行 block1 后跳出
    block1
} else if optionalStatement2; booleanExpression2 { // optionalStatement1：可选的声明语句，booleanExpression1：真假表达式，为真且上面的条件都为假则执行 block2 后跳出
    block2
} else { // 上面的条件都为假则执行 block3 后跳出
    block3
}
*/

func main() {
	// 单个 if
	if true {
		fmt.Println("if 语句")
	}

	// if ... else ...
	if false {
		fmt.Println("不会输出")
	} else {
		fmt.Println("输出")
	}

	// if ... else if ... else ...
	if 1 > 5 {
		fmt.Println("条件假，不输出")
	} else if 1 > 0 {
		fmt.Println("条件真，输出")
	} else {
		fmt.Println("不会输出")
	}

	// if else if 更多
	if v := 1; v > 5 {
		fmt.Println("v==1 假，不输出")
	} else if v++; v > 4 { // v == 2
		fmt.Println("v==2 假，不输出2")
	} else if v += 1; v == 3 { // v == 3
		fmt.Println("v==3 真，输出3")
	} else {
		fmt.Println("上面语句有真不输出")
	}

	//fmt.Println(v) // 编译报错：undefined: v，v 的作用域只在 if ... else if ... else 范围内
}
