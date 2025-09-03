package main

import (
	"fmt"
)

func main() {
	//切片就是引用类型
	sliceA := []int{1, 2, 3, 45}
	sliceB := sliceA
	sliceB[0] = 11
	fmt.Println(sliceA) // [11 2 3 45]
	fmt.Println(sliceB) // [11 2 3 45]

	//1、copy()函数复制切片
	sliceC := []int{1, 2, 3, 45}
	sliceD := make([]int, 4, 4)
	copy(sliceD, sliceC)
	fmt.Println(sliceC)
	fmt.Println(sliceD)
	sliceD[0] = 111
	fmt.Println(sliceC) // [1 2 3 45]
	fmt.Println(sliceD) // [111 2 3 45]
}
