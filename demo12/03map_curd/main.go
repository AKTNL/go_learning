package main

import "fmt"

func main() {
	//map类型数据的curd
	//1、创建 修改map类型的数据
	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["sex"] = "男"
	fmt.Println(userinfo)
	fmt.Println(userinfo["username"])

	//2、创建 修改map类型的数据
	var userinfo1 = map[string]string{
		"username": "张三",
		"age":      "20",
	}
	userinfo1["username"] = "李四"
	fmt.Println(userinfo1)

	//3、获取 查找map类型的数据
	var userinfo2 = map[string]string{
		"username": "张三",
		"age":      "20",
	}
	fmt.Println(userinfo2["age"]) // 获取

	v, ok := userinfo2["age"]
	println(v, ok)

	v1, ok1 := userinfo2["sex"]
	println(v1, ok1) // 空 和 false

	//4、删除map数据里面的key以及对应的值
	var userinfo3 = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
		"height":   "180cm",
	}
	fmt.Println(userinfo3)

	delete(userinfo3, "sex")
	fmt.Println(userinfo3)
}
