package main

import "fmt"

/*
类型		名称					存储空间		值范围
byte	字符型，unit8 别名	8-bit	表示 UTF-8 字符串的单个字节的值，对应 ASCII 码的字符值
rune	字符型，uint32 别名	32-bit	表示 单个 Unicode 字符
*/

func main() {
	var b byte  		// uint8 别名
	var r1, r2 rune 	// uint16 别名
	fmt.Printf("b: %v, r1: %v, r2: %v\n", b, r1, r2) // 默认值为0
	b = 'a'
	r1 = 'b'
	r2 = '字'
	fmt.Printf("b: %v, r1: %v, r2: %v\n", b, r1, r2) // 输出：b: 97(ASCII表示的数), r1: 98(utf-8表示的数), r2: 23383 (utf-8表示的数)

}