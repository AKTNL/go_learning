package main

import "fmt"

func main() {
	//切片的声明 初始化
	var arr1 []int
	fmt.Printf("%v - %T - 长度：%v\n", arr1, arr1, len(arr1))

	var arr2 = []int{1, 2, 3, 4}
	fmt.Printf("%v - %T - 长度：%v\n", arr2, arr2, len(arr2))

	var arr3 = []int{1: 2, 2: 4, 5: 6, 4}
	fmt.Printf("%v - %T - 长度：%v\n", arr3, arr3, len(arr3))

	fmt.Println(arr1)        // []
	fmt.Println(arr1 == nil) // true golang中申明切片以后 切片的默认值就是nil

	fmt.Println(arr2)
	fmt.Println(arr2 == nil) // false
}
