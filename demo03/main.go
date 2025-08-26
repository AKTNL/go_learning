package main

import "fmt"

func getUserinfo() (string, int) {
	return "zhangsan", 10
}

var g = "全局变量"

func main() {
	fmt.Println("声明变量")
	fmt.Println(g)
	/*
		1.var 声明变量
		2.var 变量名称 类型
		3.变量名称的命名，字母、数字、下划线，不能数字开头、关键字、保留字不能用作变量名（和其他语言一样
	*/
	/*
		var username string
		fmt.Println(username) 没有初始化，值为空

		var a1 = "张三"
		fmt.Println(a1)

		var m_ = "lisi"
		fmt.Println(m_)
	*/

	//变量定义、初始化
	/*
		var username string
		username = "zhangsan"

		fmt.Println(username)

		var username1 string = "lisi"
		fmt.Println(username1)

		var username2 = "wangwu"
		fmt.Println(username2)

		var age = 20
		var sex = "male"
		fmt.Println(username, age, sex)

		username = "zz" //赋值
		fmt.Println(username)
	*/

	/*
		一次声明多个变量
		var 变量名称，变量名称，类型
		var(
			变量名称 类型
			变量名称 类型
		)
	*/
	var a1, a2 string
	a1 = "aaa"
	a2 = "a"
	fmt.Println(a1, a2)

	var (
		username string
		age      int
		sex      string
	)
	username = "zhangsan"
	age = 20
	sex = "male"
	fmt.Println(username, age, sex)

	var (
		user_name string = "zhang"
		age1             = 21
	)
	fmt.Println(user_name, age1)

	/*
		短变量声明：只能局部，不能全局
	*/
	user := "zzz"
	fmt.Println(user)
	fmt.Printf("类型%T\n", user)

	//声明多个变量
	a, b, c := 1, 2, "C"
	fmt.Println(a, b, c)
	fmt.Printf("a类型：%T，b类型：%T,c类型：%T\n", a, b, c)

	/*
		匿名变量,忽略一个值用_
	*/
	//var u, _age = getUserinfo()
	//fmt.Println(u, _age) //zhangsan 10
	var u, _ = getUserinfo()
	fmt.Println(u) //zhangsan
	var _, ge = getUserinfo()
	fmt.Println(ge) //10
}
