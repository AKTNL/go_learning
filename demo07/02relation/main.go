package main

import "fmt"

func main() {
	var a1 = 9
	var a2 = 8
	fmt.Println(a1 == a2) // false
	fmt.Println(a1 != a2) // true
	fmt.Println(a1 > a2)  // true
	fmt.Println(a1 >= a2) // true
	fmt.Println(a1 < a2)  // false
	fmt.Println(a1 <= a2) // false
	flag := a1 > a2
	fmt.Println("flag = ", flag)

	if flag {
		fmt.Println("a1 > a2")
	}

	if a1 == a2 {
		fmt.Println("a1=a2")
	} else {
		fmt.Println("a1!=a2")
	}
}
