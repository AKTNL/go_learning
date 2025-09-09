package main

import "fmt"

// 求两个数的和
func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

// 求两个数的差
func subFn(x int, y int) int {
	fmt.Println(x, y)
	sub := x - y
	return sub
}

// 函数参数的简写
func subFn1(x, y int) int {
	sub := x - y
	return sub
}

// 函数的可变参数
func sumFn1(x ...int) int {
	fmt.Printf("%v -- %T\n", x, x) // [12 34 45 46]--[]int
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func sumFn2(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func main() {
	sum1 := sumFn(12, 3)
	fmt.Println(sum1) // 15

	sum2 := sumFn(15, 5)
	fmt.Println(sum2) // 20

	a := 20
	b := 2
	sub1 := subFn(a, b)
	fmt.Println(sub1) // 18

	sub2 := subFn1(22, 4)
	fmt.Println(sub2) // 18

	sum3 := sumFn1(12, 34, 45, 46)
	fmt.Println(sum3) // 137

	sum4 := sumFn1(1, 2, 3, 4, 5)
	fmt.Println(sum4) // 15

	sum5 := sumFn2(100, 1, 2, 3, 4)
	fmt.Println(sum5)
}
