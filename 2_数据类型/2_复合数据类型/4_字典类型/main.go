package main

import (
	"fmt"
	"unsafe"
)

/*
map 是一种键值对的无序集合，与 slice 类似也是一个引用类型。map 本身其实是个指针，指向内存中的某个空间。
声明方式与数组类似，
	声明方式：var 变量名 map[key类型]值类型
或
	使用 make 函数初始化：make(map[key类型]值类型, 初始空间大小)
其中key值可以是任何可以用==判断的值类型，对应的值类型没有要求(但定义以后只能是一种类型)
*/

func main() {
	var v1, v2 bool         // 声明变量，默认值为 false
	v1 = true               // 赋值
	v3, v4 := false, true   // 声明并赋值(更简单的初始化方式)

	fmt.Println("v1:", v1)   	// v1 输出 true
	fmt.Println("v2:", v2) 	// v2 没有重新赋值，显示默认值：false
	fmt.Println("v3:", v3) 	// v3 false// 声明后赋值
	var m map[int]string	//声明名为 m 的 map 结构, key 为 int 类型,值为 string 类型
	fmt.Println(m) 		// 输出空的map：map[]
	//m[1] = `aa`    	// 向未初始化的map中赋值报错：panic: assignment to entry in nil map

	// 声明并初始化，初始化使用{} 或 make 函数（创建类型并分配空间）
	var m1 = map[string]int{}
	var m2 = make(map[string]int)
	m1[`a`] = 11
	m2[`b`] = 22
	fmt.Println(m1) // 输出：map[a:11]
	fmt.Println(m2) // 输出：map[b:22]

	// 初始化多个值
	var m3 = map[string]string{"a": "aaa", "b": "bbb"}
	m3["c"] = "ccc"	// 增加(或修改)一个键值对
	fmt.Println(m3) // 输出：map[a:aaa b:bbb c:ccc]

	// 删除 map 中的值
	delete(m3, "a") // 删除键 a 对应的值
	fmt.Println(m3) // 输出：map[b:bbb c:ccc]

	// 查找 map 中的元素
	v, ok := m3["b"] 	//获取 m3 中 key 为"b"的值以及是否可以获取到(能否获取结果在 ok 中)
	if ok {
		fmt.Println(ok)
		fmt.Println("m3中b的值为：", v) // 输出：m3中b的值为： bbb
	}
	// 或者
	if v, ok := m3["b"]; ok { // 流程处理后面讲
		fmt.Println("m3中b的值为：", v) // 输出：m3中b的值为： bbb
	}

	fmt.Println(m3["c"]) 	// 直接取值，输出：ccc

	// map 中的值可以是任意类型(下面为 int ***声明时规定了类型)
	m4 := make(map[string][5]int)
	m4["a"] = [5]int{1, 2, 3, 4, 5}
	m4["b"] = [5]int{11, 22, 33}   	// 定义的长度为 5 没有声明的默认为 0
	fmt.Println(m4)                	// 输出：map[a:[1 2 3 4 5] b:[11 22 33 0 0]]
	// unsafe.Sizeof(m4) 用于过去 m4 所占用的内存空间
	fmt.Println(unsafe.Sizeof(m4)) 	// 输出：8，为8个字节，map其实是个指针，指向某个内存空间
	fmt.Println("v4:", v4) 	// v4 true
}