package main

import (
	"fmt"
	"strings"
)

func main() {
	// 统计一个字符串中每个单词出现的任务
	var str = "how do you do how do you do"

	var strSlice = strings.Split(str, " ")
	fmt.Println(strSlice)

	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)
}
