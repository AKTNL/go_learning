package main

import "fmt"

func main() {
	var a = 6
	var b = 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)

	var c = a * b
	fmt.Println(c)

	//1、除法注意：如果运算的数都是整数，那么除后去掉小数部分，保留整数部分
	a = 10
	b = 3
	fmt.Println(a / b) // 3

	var a1 = 10.0
	var b1 = 3.0
	fmt.Println(a1 / b1) // 3.333333333333333335

	//2、取余注意 余数=被除数-（被除数/除数）*除数
	a = 10
	b = 3
	fmt.Println(a % b) // 1

	fmt.Println(-10 % 3) // -10 - (-10 / 3) * 3 = -1

	fmt.Println(10 % -3) // 10 - (10 / -3) * -3 = 1

	//1、注意：在golang中，++和--只能独立使用
	// a = i++ 错误，i++只能独立使用

	//2、注意：在golang中没有前++和--
	// --a 错误写法

	//3、正确写法
	var a4 = 12
	a4++
	fmt.Println(a4)

}
