package main

import "fmt"

/*
顾名思义让程序跳到某个指定行重新执行。
continue 语句用于跳过本次循环继续下次循环
break 用于结束循环 	他两个也可以配合标签来使用跳出多重循环
goto 跳转到指定标签位置，标签区分大小写,每一个标签代表一个代码块。--> 一般会把 标签 写在goto 跳转的后面,最大程度上避免死循环
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

LABL1:
	for i := 0; i < 2; i++ {
		for i := 0; i < 10; i++ {
			fmt.Print(i, " ")
			if i > 4 {
				break LABL1 // 跳出外层循环(通过标签跳出多层)
				//continue LABL1
			}
		}
	}
	fmt.Println()

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
