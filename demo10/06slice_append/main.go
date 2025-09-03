package main

import "fmt"

func main() {
	//1、append（）方法的使用
	var sliceA []int
	fmt.Printf("%v -- %v -- %v\n", sliceA, len(sliceA), cap(sliceA)) // [] -- 0 -- 0
	sliceA = append(sliceA, 12, 23, 35, 465)
	//sliceA = append(sliceA, 24)
	fmt.Printf("%v -- %v -- %v\n", sliceA, len(sliceA), cap(sliceA)) // [12 23 35 465] -- 4 -- 4

	//2、append()方法还可以合并切片
	sliceB := []string{"php", "java"}
	sliceC := []string{"nodejs", "go"}
	sliceB = append(sliceB, sliceC...)
	fmt.Println(sliceB)

	//3、切片的扩容策略
	var sliceD []int
	for i := 1; i <= 10; i++ {
		sliceD = append(sliceD, i)
		fmt.Printf("%v -- 长度%d -- 容量%d\n", sliceD, len(sliceD), cap(sliceD))
	}
}
