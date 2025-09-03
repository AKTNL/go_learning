package main

import "fmt"

func main() {
	var sliceA = make([]int, 4, 8)
	fmt.Println(sliceA) // [0 0 0 0]
	sliceA[0] = 10
	sliceA[1] = 12
	sliceA[2] = 40
	sliceA[3] = 30
	fmt.Println(sliceA)

	sliceB := []string{"php", "java", "go"}
	sliceB[2] = "golang"
	fmt.Println(sliceB)

	//golang中没法通过下标的方式给切片扩容 append()
	var sliceC []int
	fmt.Printf("%v -- %v -- %v\n", sliceC, len(sliceC), cap(sliceC)) // [] -- 0 -- 0
	//sliceC[0] = 1 报错：下标越界了
	fmt.Println(sliceC)

	//golang中给切片扩容的话要用到append()方法
}
