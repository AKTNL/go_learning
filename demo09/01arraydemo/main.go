package main

import (
	"fmt"
)

func main() {

	//1、数组的长度是类型的一部分
	var arr1 [3]int
	var arr2 [4]int
	var strArr [3]string

	fmt.Printf("arr1:%T  arr2:%T  strArr:%T\n", arr1, arr2, strArr)

	//2、数组的初始化 第一种方法
	fmt.Println(arr1, strArr)

	arr1[0] = 23
	arr1[1] = 10
	arr1[2] = 24
	fmt.Println(arr1)

	strArr[0] = "php"
	strArr[1] = "java"
	strArr[2] = "golang"
	fmt.Println(strArr)

	//2、数组的初始化 第二种方法
	var arr11 = [3]int{23, 334, 5}
	fmt.Println(arr11)
	var arr12 = [3]string{"php", "nodejs", "golang"}
	fmt.Println(arr12)

	arr13 := [3]string{"php", "nodejs", "golang"}
	fmt.Println(arr13)

	//3、数组的初始化 第三种方法
	var arr_1 = [...]int{12, 345, 67, 89, 2456}
	fmt.Println(arr_1)
	fmt.Println(len(arr_1)) // len()打印数组的长度
	arr_2 := [...]string{"php", "nodejs", "js", "golang"}
	fmt.Println(arr_2)

	//改变数组里面的值
	arr_2[0] = "phper"
	fmt.Println(arr_2)

	//4、数组的初始化 第四种方法
	arr := [...]int{0: 1, 1: 10, 2: 20, 5: 50}
	fmt.Println(len(arr))
	fmt.Println(arr)

	//5、数组的循环遍历
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for i := 0; i < len(arr_2); i++ {
		fmt.Println(arr_2[i])
	}

	for k, v := range arr_2 {
		fmt.Printf("key:%v value:%v\n", k, v)
	}
}
