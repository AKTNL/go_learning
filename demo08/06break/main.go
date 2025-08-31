package main

import "fmt"

func main() {

	//1、表示当i=2的时候跳出当前循环
	for i := 1; i <= 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("继续执行")

	//2、表示当i=2的时候跳出当前循环
	/*
		i=0
			i=0 j=0
			i=0 j=1
			i=0 j=2
		i=1
			i=1 j=0
			i=1 j=1
			i=1 j=2
	*/
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}

	//2、break在switch（开关语句）中执行一条case后跳出语句的作用

	//3、在多重循环中，可以用标号label标想出break的循环
lable1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break lable1
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}
}
