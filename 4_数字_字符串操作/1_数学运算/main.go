package main

import (
	"fmt"
	"reflect"
)
/*
Go 语言具有严格的静态类型限制，运算符操作的数字类型必须是相同类型数据。
	且数字操作不能超出该类型的取值范围，不然计算溢出得到的结果就是错的
本节主要讲述,基本的数值运算:加减乘除,取模
*/
func main() {
	fmt.Println("---------加 + ------------")
	var v1 uint8 = 1
	var v2 uint8 = 2
	u1 := v1 + v2
	fmt.Println(u1)
	u1 += 255
	fmt.Println(u1) // 溢出后计算的值就是错的，输出：2

	fmt.Println("---------减 - ------------")
	u2 := v2 - v1
	fmt.Println(u2)

	fmt.Println("---------乘 * ------------")
	var f1 float32 = 1.2
	var f2 float32 = 0.2
	fmt.Println(f1 * f2)

	fmt.Println("---------除 / ------------")
	// 小数会得到小数结果
	fmt.Println(f2 / f1)
	fmt.Println(reflect.TypeOf(f1 / f2))

	// 默认整数类型运算时会返回整数类型(取整数部分)
	n1 := 11
	n2 := 3
	fmt.Println(n1 / n2) // 输出：3，小数被丢弃

	// 求模、取余 %
	fmt.Println(n1 % n2) // 输出余数：2

	n3 :=
}
