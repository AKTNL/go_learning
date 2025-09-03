package main

import "fmt"

func main() {
	s2 := "你好golang"
	runeStr := []rune(s2)
	fmt.Println(runeStr) // [20320 22909 103 111 108 97 110 103]

	runeStr[0] = '大'
	fmt.Println(string(runeStr))
}
