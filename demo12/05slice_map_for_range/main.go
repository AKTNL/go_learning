package main

import "fmt"

func main() {
	var userinfo = make([]map[string]string, 3, 3)
	// fmt.Println(userinfo[0]) // map[] nil
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
	if userinfo[2] == nil {
		userinfo[2] = make(map[string]string)
		userinfo[2]["username"] = "王五"
		userinfo[2]["age"] = "21"
		userinfo[2]["height"] = "175cm"
	}
	fmt.Println(userinfo)
	for _, v := range userinfo {
		//fmt.Println(v)
		for key, value := range v {
			fmt.Printf("kay:%v -- value:%v\n", key, value)
		}
	}

}
