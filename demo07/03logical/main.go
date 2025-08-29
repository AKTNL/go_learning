package main

import "fmt"

// 定义一个方法
func test() bool {
	fmt.Println("test...")
	return true
}

func main() {
	var a = 23
	var b = 8
	fmt.Println(a > 10 && b < 10) // true
	fmt.Println(a > 24 && b < 10) // false
	fmt.Println(a > 5 && b < 6)   // false
	fmt.Println(a == 5 || b < 6)  // false

	fmt.Println(a > 10 || b < 10) // true
	fmt.Println(a > 24 || b < 10) // true
	fmt.Println(a > 5 || b < 6)   // true
	fmt.Println(a == 5 || b < 6)  // false

	flag := true
	fmt.Println(!flag)

	//逻辑与和逻辑或短路

	/*
		test...
		执行
	*/
	a = 10
	if a > 9 && test() {
		fmt.Println("执行")
	}

	/*
		输出：
		空
	*/
	if a > 11 && test() {
		fmt.Println("执行")
	}

	/*
		输出：
		test...
		执行
	*/
	if a > 11 || test() {
		fmt.Println("执行")
	}

	/*
		输出：
		执行
	*/
	if a < 11 || test() {
		fmt.Println("执行")
	}
}
