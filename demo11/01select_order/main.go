package main

import "fmt"

func main() {
	var numSlice = []int{9, 8, 7, 6, 5, 4}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] > numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}

	}
	fmt.Println(numSlice)
}
