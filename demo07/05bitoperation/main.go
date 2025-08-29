package main

import "fmt"

func main() {
	var a = 5                      // 101
	var b = 2                      // 010
	fmt.Println("a & b = ", a&b)   // 000 值：0
	fmt.Println("a | b = ", a|b)   // 111 值：7
	fmt.Println("a ^ b = ", a^b)   // 111 值：7
	fmt.Println("a << b = ", a<<b) // 10100 值：20 a左移b位 5*2的2次方
	fmt.Println("a >> b = ", a>>b) // 1 值：1 右移 a/2的b次方
}
