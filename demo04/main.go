package main

import (
	"fmt"
)

func main() {
	// var username = "zhangsan"
	// username = "lisi"
	// username = "wangwu"
	// fmt.Println(username)

	//常量:常量的值不可以改变
	const pi = 3.14159
	fmt.Println(pi)

	//多个常量也可以一起声明
	const (
		a = "A"
		b = "B"
	)
	fmt.Println(a, b)

	//可以同时声明多个常量，如果省略值则和上面一样
	const (
		n = 100
		n1
		n2
	)
	fmt.Println(n, n1, n2)

	//iota,计数器和const一起使用
	const A = iota
	fmt.Println(A)

	const (
		a1 = iota
		a2
		a3
		a4
		_
		a6
	)
	fmt.Println(a1, a2, a3, a4, a6)

	const (
		b1, b2 = iota + 1, iota + 2
		b3     = 100
		b4     = iota
		b5
	)
	fmt.Println(b1, b2, b3, b4, b5)

	const (
		c1, c2 = iota + 1, iota + 2 //1,2
		c3, c4                      //2,3
		c5, c6                      //3,4
	)
	fmt.Println(c1, c2, c3, c4, c5, c6)

	//定义变量
	d1, d2 := 20, 30
	fmt.Println(d1, d2)

	//变量名字区分大小写
	var age = 20
	var Age = 30
	fmt.Println(age, Age)

	var maxAge = 20
	var MaxAge = 30
	fmt.Println(maxAge, MaxAge)

	var DNS = "192.112.23.2"
	fmt.Println(DNS)
}
