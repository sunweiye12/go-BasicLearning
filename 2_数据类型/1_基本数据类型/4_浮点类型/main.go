package main
import "fmt"
/*
类型		名称				存储空间		值范围							数据级别
float32	32位浮点数		32-bit		IEEE-754 1.4e-45 ~ 3.4e+38		精度6位小数
float64	64位浮点数		64-bit		IEEE-754 4.94e-324 ~ 1.798e+308 	精度15位小数
*/
func main() {
	// 浮点型，f32精度6位小数，f64位精度15位小数
	var f32 float32
	var f64 float64
	fmt.Printf("f32: %f, f64: %f\n", f32, f64) // 默认值都为 0.000000
	f32 = 1.123456789
	f64 = 1.12345678901234567
	fmt.Printf("f32: %v, f64: %v\n", f32, f64) // 末位四舍五入，输出：f32: 1.1234568, f64: 1.1234567890123457
}