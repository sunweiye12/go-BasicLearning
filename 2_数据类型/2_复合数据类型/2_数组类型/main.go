package main

import "fmt"
/*
数组为一组相同数据类型数据的集合**，数组定义后大小固定**，不能更改，每个元素称为element，
声明的数组元素默认值都是对应类型的0值。
而且数组在Go语言中是一个值类型（value type），
所有值类型变量在赋值和作为参数传递时都会产生一次复制动作，即对原值的拷贝。(后续的操作不会改变原值)
*/

func main() {
	// 1.声明后赋值
	// var <数组名称> [<数组长度>]<数组元素类型>
	var arr [2]int   // 数组元素的默认值都是 0
	fmt.Println(arr) // 输出：[0 0]
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr) // 输出：[1 2]

	// 2.声明并赋值
	// var <数组名称> = [<数组长度>]<数组元素>{元素1,元素2,...}
	var intArr = [2]int{1, 2}
	strArr := [3]string{"aa", `bb`, `cc`}	//字符串用单引号(或双引号)括起来
	fmt.Println(intArr) // 输出：[1 2]
	fmt.Println(strArr) // 输出：[aa bb cc]

	// 3.声明时不设定大小，赋值后语言本身会计算数组大小
	// var <数组名称> [<数组长度>]<数组元素> = [...]<元素类型>{元素1,元素2,...}
	var arr1 = [...]int{1, 2}
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr1) // 输出：[1 2]
	fmt.Println(arr2) // 输出：[1 2 3]
	//arr1[2] = 3 // 编译报错，数组大小已设定为2

	// 4.声明时不设定大小，赋值时指定索引
	// var <数组名称> [<数组长度>]<数组元素> = [...]<元素类型>{索引1:元素1,索引2:元素2,...}
	var arr3 = [...]int{1: 22, 0: 11, 2: 33}
	arr4 := [...]string{2: "cc", 1: "bb", 0: "aa"}
	fmt.Println(arr3) // 输出：[11 22 33]
	fmt.Println(arr4) // 输出：[aa bb cc]

	// 遍历数组
	for i := 0; i < len(arr4); i++ {
		v := arr4[i]
		fmt.Printf("i:%d, value:%s\n", i, v)
	}
}