package main

import "fmt"

func main() {
	var a = 20
	a = 21
	fmt.Println(a)

	a = 23 + 2
	fmt.Println(a)

	a = 10
	b := a
	fmt.Println(b)

	b = a + 2
	fmt.Println(b)

	a += 3 // 等价于 a = a + 3
	fmt.Println(a)

	a = a + 3
	fmt.Println(a)

	a = 10
	a -= 3 // a=a-3
	fmt.Println(a)

	a *= 3 // a=a*3
	fmt.Println(a)

	a = 10
	a /= 3         //a=a/3
	fmt.Println(a) // 结果3

	var a1 float64 = 10
	a1 /= 3
	fmt.Println(a1) //结果：3.33333333333335

	a = 10
	a %= 3         // a=a%3
	fmt.Println(a) //结果：1
}
