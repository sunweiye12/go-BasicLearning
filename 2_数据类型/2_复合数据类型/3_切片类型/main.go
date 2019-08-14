package main

import "fmt"

/*
因为数组的长度定义后不可修改，所以需要切片来处理可变长数组数据。
	切片可以看作是一个可变长的数组，是一个引用类型。
它包含三个数据：1.指向原生数组的指针，2.切片中的元素个数，3.切片已分配的存储空间大小
声明一个切片，或从数组中取一段作为切片数据
*****
使用make直接创建切片，语法：make([]类型, 大小，预留空间大小)，
make() 函数用于声明slice切片、map字典、channel通道。
*/

func main() {
	var sl []int                // 声明一个切片
	sl = append(sl, 1, 2, 3, 4) // 往切片中追加值
	fmt.Println(sl)             // 输出：[1 2 3 4]

	var arr = [5]int{1, 2, 3, 4, 5} // 初始化一个数组
	var sl1 = arr[0:2]              // 冒号:左边为起始位（包含起始位数据），右边为结束位（不包含结束位数据）；不填则默认为头或尾
	var sl2 = arr[3:]
	var sl3 = arr[:5]

	fmt.Println(sl1) // 输出：[1 2]
	fmt.Println(sl2) // 输出：[4 5]
	fmt.Println(sl3) // 输出：[1 2 3 4 5]

	// -----------------append()方法的作用----------------
	sl1 = append(sl1, 11, 22) // 追加元素
	fmt.Println(sl1)          // 输出：[1 2 11 22]

	// -----------------copy()方法的作用------------------
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{9, 8, 7}
	//copy(s2,s1)	// 将 s2 赋值给 s1  --> 从前往后来以此 copy知道任意一个结束,不会改变原来切片的大小
	copy(s2[1:], s1[4:])
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("-------------通过 make 来创建切片---------------")
	//通过 make 来创建切片
	makeSlice()
}

//通过 make 创建 slice 实现
func makeSlice() {
	var sl1 = make([]int, 5) // 定义元素个数为5的切片
	// 切片类型底层为数组结构,因为重新分配内存是开销比较大的操作,因此我们知道大体数值量的情况下可以提前开辟足够的缓存空间
	sl2 := make([]int, 5, 10)         // 定义元素个数5的切片，并预留10个元素的存储空间（相当于缓存空间为 10 个） 当超出缓存以后会进行 2 倍扩容
	fmt.Println(len(sl2), cap(sl2))   // len 获取长度  cap 获取总容量(缓存)
	sl3 := []string{`aa`, `bb`, `cc`} // 直接创建并初始化包含3个元素的数组切片

	fmt.Println(sl1, len(sl1)) // 输出：[0 0 0 0 0] 5
	fmt.Println(sl2, len(sl2)) // 输出：[0 0 0 0 0] 5
	fmt.Println(sl3, len(sl3)) // [aa bb cc] 3

	sl1[1] = 1 // 声明或初始化大小中的数据，可以指定赋值
	sl1[4] = 4
	//sl1[5] = 5 // 编译报错，超出定义大小
	sl1 = append(sl1, 5)       // 可以追加元素
	fmt.Println(sl1, len(sl1)) // 输出：[0 1 0 0 4 5] 6

	sl2[1] = 1
	// append 只能用于切片元素的拼接
	sl2 = append(sl2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(sl2, len(sl2)) // 输出：[0 1 0 0 0 1 2 3 4 5 6 7 8 9 10 11] 16

	// 遍历切片
	size := len(sl2)
	for i := 0; i < size; i++ {
		v := sl2[i] //通过下标取到值
		fmt.Printf("i: %d, value:%d \n", i, v)
	}
}
