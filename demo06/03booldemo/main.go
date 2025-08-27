package main

import "fmt"

func main() {
	/*
		默认值为false
		不允许将整型强制转换为布尔型
		无法参与数值运算，无法与其他类型转换
	*/
	var flag = true
	fmt.Printf("%v--%T\n", flag, flag)

	//1、布尔变量的默认值为false
	var b bool
	fmt.Printf("%v--%T\n", b, b)

	//2、string型变量的默认值为空
	var s string
	fmt.Printf("%v--%T\n", s, s)

	//3、int型变量的默认值为0
	var i int
	fmt.Printf("%v--%T\n", i, i)

	//4、float型变量的默认值为0
	var f float64
	fmt.Printf("%v--%T\n", f, f)

	//5、go语言不允许将整型转换为布尔型
	var a = true //注意：a = 1 或者 其他数据类型 的话会报错
	if a {
		fmt.Printf("true\n")
	}

}
