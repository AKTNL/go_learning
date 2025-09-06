package main

import "fmt"

func main() {
	// var userinfo = []string{"张三", "李四"}
	var userinfo = make([]map[string]string, 3, 3)
	fmt.Println(userinfo[0]) // map[] nil
	fmt.Println(userinfo[0] == nil)
	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["height"] = "180cm"
	}
	if userinfo[1] == nil {
		userinfo[1] = make(map[string]string)
		userinfo[1]["username"] = "李四"
		userinfo[1]["age"] = "22"
		userinfo[1]["height"] = "170cm"
	}
	fmt.Println(userinfo)
}
