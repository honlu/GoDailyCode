/*
面试题：
1. 实现string转为float64
2. 实现float64小数位6位转为2位
3. 常见数据类型的互相转换

ParseFloat 将字符串 s 转换为具有 bitSize 指定精度的浮点数：对于 float32 为 32，对于 float64 为 64。
当 bitSize=32 时，结果仍然是 float64 类型，但它可以转换为 float32 而不会改变它的值。
ParseFloat 接受十进制和十六进制浮点数语法。如果 s 格式正确且接近有效浮点数，ParseFloat 返回使用 IEEE754 无偏舍入舍入的最接近的浮点数。
 (仅当十六进制表示中的位数多于尾数中的位数时，才对十六进制浮点值进行四舍五入。)
https://vimsky.com/examples/usage/golang_strconv_ParseFloat.html

关于下面的具体类型，都可以换成接口变量，i.(Type)的方式进行类型断言再转换。
*/
package main

import (
	"fmt"
	"strconv"
)

// string保存的是float64数
func stringToFloat64(s string) {
	num, err := strconv.ParseFloat(s, 64) // 64在这里指
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T, %v\n", num, num)
}

/*
小数位将减少
*/
func Decimal(num float64) {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	fmt.Printf("%T, %v\n", num, num)
}

// 常见类型互相转换
func xToY() {
	// int 转string
	s0 := 1
	str := strconv.Itoa(s0)
	fmt.Printf("int to string: %T, %v\n", str, str)

	// string 转 int
	s1 := "123"
	num, _ := strconv.Atoi(s1)
	fmt.Printf("string to int: %T, %v\n", num, num)

	// int32转string, 不能直接转为string,再使用strconv.Itoa时要先转为int；或者用FormatInt需要先转为int64，strconv.FormatInt(int64(i), 10)
	s2 := int32(123)
	str = strconv.Itoa(int(s2))
	fmt.Printf("int32 to string: %T, %v\n", str, str)

	// string转int32, string无法直接转为int32，只能先转为int,再转为int32
	s3 := "456"
	t, _ := strconv.Atoi(s3)
	t3 := int32(t)
	fmt.Printf("string to int32: %T, %v\n", t3, t3)

	// int64转string
	s4 := int64(123456)
	str = strconv.FormatInt(s4, 10)
	fmt.Printf("int64 to string: %T, %v\n", str, str)

	// string转int64
	s5 := "123456789"
	t5, _ := strconv.ParseInt(s5, 10, 64)
	fmt.Printf("string to int64: %T, %v\n", t5, t5)

	// float32转string
	s6 := float32(123.12)
	str = fmt.Sprintf("%f", s6)
	fmt.Printf("float32 to string: %T, %v\n", str, str)

	// string转float32
	// string无法直接转为float32，只能先转化为float64，再转为float32
	s7 := "7123456"
	num64, _ := strconv.ParseFloat(s7, 32)
	fmt.Printf("string to float64: %T, %v\n", num64, num64)
	t7 := float32(num64)
	fmt.Printf("float64 to float32 : %T, %v\n", t7, t7)

	// float64转string
	s8 := 456.123
	str = strconv.FormatFloat(s8, 'f', -1, 64) // FormatFloat 根据格式 fmt 和精度 prec 将浮点数 f 转换为字符串。
	// 假设原始值是从 bitSize 位的浮点值(32 表示 float32，64 表示 float64)获得的，它会对结果进行四舍五入。
	fmt.Printf("float64 to string: %T, %v\n", str, str)

	// string转float64
	s9 := "45646554.456464"
	t9, _ := strconv.ParseFloat(s9, 64)
	fmt.Printf("string to float64: %T, %v\n", t9, t9)
}

func main() {
	// string to float64
	s0 := "3.456789"
	stringToFloat64(s0) // float64, 3.456789

	// float64 6位小数转2位小数
	s1 := 3.456789
	Decimal(s1) //  float64, 3.456789

	// 常见类型转换
	xToY()
}
