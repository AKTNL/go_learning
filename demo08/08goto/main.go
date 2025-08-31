package main

import "fmt"

func main() {
	var n = 30
	if n > 24 {
		fmt.Println("成年人")
		goto lable3
	}
	fmt.Println("aaa")
lable3:
	fmt.Println("bbb")
}
