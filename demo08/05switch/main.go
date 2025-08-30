package main

import "fmt"

func main() {
	/*
		switch expression{
		case condition:
		}
	*/

	var ext = ".html"
	if ext == ".html" {
		fmt.Println("text/html")
	} else if ext == ".css" {
		fmt.Println("text/css")
	} else if ext == ".js" {
		fmt.Println("text/javascript")
	} else {
		fmt.Println("找不到此后缀")
	}

	//1、swith case的基本使用
	var extname = ".js"
	switch extname {
	case ".html":
		fmt.Println("text/html")
		break
	case ".css":
		fmt.Println("text/css")
		break
	case ".js":
		fmt.Println("text/javascript")
		break
	default:
		fmt.Println("找不到此后缀")
	}

	//2、switch case的另一种写法
	switch extname := ".css"; extname {
	case ".html":
		fmt.Println("text/html")
		break
	case ".css":
		fmt.Println("text/css")
		break
	case ".js":
		fmt.Println("text/javascript")
		break
	default:
		fmt.Println("找不到此后缀")
	}

	//3、一个分支可以有多个值，多个case值中间使用英文逗号分隔
	//判断一个数是不是偶数
	var n = 8
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数") //golang中break可以写也可以不写
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}

	var score = "A" //ABC及格
	switch score {
	case "A", "B", "C":
		fmt.Println("及格") //golang中break可以写也可以不写
	case "D":
		fmt.Println("不及格")
	}

	switch score := "D"; score {
	case "A", "B", "C":
		fmt.Println("及格") //golang中break可以写也可以不写
	case "D":
		fmt.Println("不及格")
	}

	//4、分支可以使用表达式
	var age = 30
	switch {
	case age < 24:
		fmt.Println("好好学习")
	case age >= 24 && age <= 60:
		fmt.Println("好好工作")
		fallthrough
	case age > 60:
		fmt.Println("注意身体")
	default:
		fmt.Println("输入错误")
	}

	//5、 switch 的穿透 fallthrough
	//可以执行满足条件的case的下一条

}
