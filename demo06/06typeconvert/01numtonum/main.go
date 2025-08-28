package main

import "fmt"

func main() {
	//1、整型和整型之间的转换
	var a int8 = 20
	var b int16 = 40
	fmt.Println(int16(a) + b)

	//2、浮点型和浮点型之间的转换
	var a1 float32 = 20.2
	var b1 float64 = 40.3
	fmt.Println(float64(a1) + b1)

	//3、整型和浮点型之间的转换
	var a2 float32 = 20.23
	var b2 int = 40
	fmt.Println(a2 + float32(b2))

	//注意：转换的时候建议从地位转到高位
}
