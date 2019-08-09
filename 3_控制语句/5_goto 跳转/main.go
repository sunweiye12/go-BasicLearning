package main

import "fmt"

/*
顾名思义让程序跳到某个指定行重新执行。
使用 goto 语句必须在当前函数内定义跳转标签，标签区分大小写,每一个标签代表一个代码块。
程序中一般不建议使用 goto 语句，过多的 goto 语句会打乱程序逻辑难以控制。
语法:
func funcName() {
    ... //
    GotoTag: // goto 跳转标签
    ...
    goto GotoTag // 跳转到 GotoTag 行重新执行
    ...
}
*/
func main() {

	i := 0

	// i 递增输出
LOOP123:
	for i < 10 {
		i++
		if i == 5 { // i = 5 时跳出 for 循环，不会打印 5
			goto LOOP123 // 跳到指定的标签处,每一个标签都是一个代码块
		}
		fmt.Println(i)
	}

	// 递减输出 i 到 0
TAG:
	fmt.Println(i)
	i--
	if i < 0 {
		goto TAG
	}
	//fmt.Println("end")

}
