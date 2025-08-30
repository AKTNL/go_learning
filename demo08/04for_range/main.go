package main

import "fmt"

func main() {
	var str = "你好golang"
	for k, v := range str {
		fmt.Printf("key = %v val = %c\n", k, v)
	}
	var arr = []string{"php", "java", "node", "golang"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, v := range arr {
		fmt.Println(v)
	}
}
