package main

import (
	"fmt"
)

/*
for 循环在 Go 语言中有多种格式选择：直接看例子
*/
func main() {

	//输出 0-10 的数字(类似于 while 的实现)
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
	fmt.Println("-----------------------------")
	//可以当做 while 来用
	v := 1
	for v <= 10 {
		fmt.Print(v, " ")
		v++
	}

	fmt.Println()
	fmt.Println("--------------遍历字符串---------------")
	str := "hello 世界"
	//utf-8 遍历
	for i := 0; i < len(str); i++ { // len()函数用于获取字符串长度
		chr := str[i]
		//caye := reflect.TypeOf(chr)
		fmt.Printf("%v", chr) // 输出：uint8
	}
	//unicode 遍历
	for index, char := range str {
		fmt.Print(index, char)   // 输出字符起始位，及 Unicode 编码，如：6 19990
		fmt.Printf("%q\n", char) // 格式化输出单个字符字面值，如：'世'
	}

	fmt.Println()
	fmt.Println("--------------遍历数组,切片---------------")
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("i: %d, v: %d\n", i, v)
	}
	// 切片
	sl := arr[2:]
	for _, v := range sl { // 赋值给 '_' 的值会被忽略掉
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println("--------------遍历字典---------------")
	var m = map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	for key := range m {
		fmt.Println(key, m[key])
	}

	fmt.Println()
	fmt.Println("--------------遍历沟道---------------")
	ch := make(chan int, 2) // 创建了一个缓存为 2 的 channel 结构
	ch <- 1
	ch <- 2
	//fmt.Println( <-ch )
	for v := range ch {
		fmt.Println(v)
	}

}
