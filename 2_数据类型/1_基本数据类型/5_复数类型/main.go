package main
import "fmt"
/*
类型			名称										存储空间	值范围	数据级别
complex64	复数，含 float32 位实数和 float32 位虚数	64-bit	实数、虚数的取值范围对应 float32
complex128	复数，含 float64 位实数和 float64 位虚数	128-bit	实数、虚数的取值范围对应 float64
*/
func main() {
	// 复数型
	var c64 complex64
	var c128 complex128
	fmt.Printf("c64: %v, c128: %v\n", c64, c128) // 实数、虚数的默认值都为0
	c64 = 1.12345678 + 1.12345678i
	c128 = 2.1234567890123456 + 2.1234567890123456i
	fmt.Printf("c64: %v, c128: %v\n", c64, c128) // 输出：c64: (1.1234568+1.1234568i), c128: (2.1234567890123457+2.1234567890123457i)

	// 虚数的运算
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x*y)                 // "(-5+10i)"
	fmt.Println(real(x*y))  // 获取实数部分	// "-5"
	fmt.Println(imag(x*y))	// 获取虚数部分
}