package main

import "fmt"

func main() {
	//1、最简单的if语句
	flag := true
	if flag {
		fmt.Println("flag = true")
	}

	age := 30 //当前区域内的全局变量
	if age > 20 {
		fmt.Println("成年人")
	}

	//2、if语句的另一种写法
	if age := 34; age > 20 { // 局部变量
		fmt.Println("成年人1")
	}

	//3、探讨上面两种写法的区别

	if age > 20 {
		fmt.Println("成年人", age)
	}
	fmt.Println(age)

	//4、输入一个人的成绩，如果成绩大于等于90输出A,如果小于90大于75输出B,否则输出C
	var score = 93
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	if score := 75; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//5、if else要注意的细节
	//1、if{}不能省略
	if age > 20 {
		fmt.Println("成年人")
	}
	//2、{必须紧挨着条件
	if age > 20 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年")
	}

	//6、求两个数的最大值
	var a = 34
	var b = 24
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Println("a和b的最大值是", max)
}
