package main

import "fmt"

func main() {
	//1、练习：打印0-50所有偶数
	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//2、练习：求1+2+3+4+...+100的和
	/*
		sum += sum + 1   0 + 1
		sum += sum + 2   0 + 1 + 2
	*/
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//3、练习：打印1～100之间所有是9的倍数的整数的个数及总和
	var sum1 = 0
	var count = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			fmt.Println(i)
			sum1 += i
			count++
		}
	}
	fmt.Println(sum1, count)

	//4、练习：计算5的阶乘（12345 n的阶乘12....n)
	/*
		1、sum = 1 * 1
		2、sum = sum * 2   1 * 1 * 2
		3、sum = sum * 3   1 * 1 * 2 * 3
	*/
	var sum2 = 1
	for i := 1; i <= 5; i++ {
		sum2 *= i
	}
	fmt.Println(sum2)

	//5、练习：打印一个矩形
	/*
	****
	****
	****
	 */
	for i := 1; i <= 12; i++ {
		fmt.Print("*")
		if i%4 == 0 {
			fmt.Println("")
		}
	}

	//6、for循环的嵌套
	/*
		for循环嵌套的一个执行流程
		1、i = 0 打印5个* 一个换行
		2、i = 1 打印5个* 一个换行
		.......
		6、i = 5 跳出循环
	*/
	var row = 5
	var column = 8
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//7、练习：打印一个三角形
	/*
	*
	**
	 */
	row = 5
	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//8、练习：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Println("")
	}
}
