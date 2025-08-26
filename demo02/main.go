package main

import "fmt"

func main() {
	//fmt.Println("你好 golang！")
	//fmt.Print("hello golang")
	// fmt.Printf("hi,golang~")

	/*
		fmt.Println("你好 golang！")
		fmt.Print("hello golang")
		fmt.Printf("hi,golang~")
	*/

	/*
		fmt.Println("我是kevin") 换行，和java一样

		fmt.Print("A", "B", "C")
		fmt.Println("D", "E", "f")

		var a = "aaa" go语言中变量定义后必须要使用
		fmt.Println(a)

		fmt.Printf("%v", a)
	*/
	/*
		var a int = 10
		var b int = 3
		var c int = 5

		fmt.Println("a=", a, "b=", b, "c=", c)
		fmt.Printf("a=%v b=%v c=%v\n", a, b, c)
	*/

	//类型推导方式定义变量
	a := 10
	b := 20
	c := 30
	fmt.Printf("a=%v b=%v c=%v\n", a, b, c)

	//打印变量类型
	fmt.Printf("a=%v a的类型是%T\n", a, a)

}
