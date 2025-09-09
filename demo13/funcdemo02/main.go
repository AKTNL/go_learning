package main

import "fmt"

func sumFn(x, y int) int {
	return x + y
}

// return 关键词一次可以返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名法
func calc1(x, y int) (sum int, sub int) {
	fmt.Println(sum, sub)
	sum = x + y
	sub = x - y
	fmt.Println(sum, sub)
	return
}

func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func test() {
	fmt.Println("执行...")
}

func main() {
	sum1 := sumFn(10, 2)
	fmt.Println(sum1) // 12

	a, b := calc(10, 2)
	fmt.Println(a, b) // 12 8

	a1, b1 := calc1(10, 2)
	fmt.Println(a1, "---", b1)

	a2, b2 := calc2(10, 2)
	fmt.Println(a2, b2)

	a3, _ := calc2(10, 2)
	fmt.Println(a3)

	_, b3 := calc2(10, 2)
	fmt.Println(b3)

	test()
}
