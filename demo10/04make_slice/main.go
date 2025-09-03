package main

import "fmt"

func main() {
	// make()函数创建一个切片 make([]T,size,cap)
	var sliceA = make([]int, 4, 8)
	fmt.Println(sliceA) // [0 0 0 0]
	fmt.Printf("%d -- %d\n", len(sliceA), cap(sliceA))
}
