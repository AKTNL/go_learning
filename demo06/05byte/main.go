package main

import (
	"fmt"
)

func main() {
	//1.golang中定义字符 字符属于int类型
	var a = 'a'
	fmt.Printf("值：%v  类型：%T\n", a, a) //值97  类型int32

	//2、原样输出字符
	fmt.Printf("值：%c  类型：%T\n", a, a) //值a

	//3、定义一个字符串，输出里面的字符
	var str = "this" // 占用4个字节
	fmt.Printf("值：%v 原样输出：%c  类型:%T\n", str[2], str[2], str[2])

	//4、一个汉字占用3个字节（utf-8）一个字母占用一个字节
	/*
		unsafe.Sizeof()没法查看string类型数据所占用的存储空间
		用len()
	*/

	fmt.Println(len(str)) //4

	var s = "你好go"
	fmt.Println(len(s)) // 8

	//golang中使用uft-8编码，编码后的值就是int类型
	a = '你'
	fmt.Printf("值：%v  类型：%T\n", a, a) //Unicode编码10进制 值：20320 类型：int32
	fmt.Printf("值：%c  类型：%T\n", a, a)

	//6、通过循环输出字符串里的字符
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v (%c)  ", s[i], s[i])
	}

	for _, v := range s { //rune
		fmt.Printf("%v (%c)  ", v, v)
	}
	fmt.Println()

	//7、修改字符串
	s1 := "big"
	byteStr := []byte(s1)
	byteStr[0] = 'p'
	fmt.Println(string(byteStr))

	s2 := "你好golang"
	byteStr2 := []rune(s2)
	byteStr2[0] = '大'
	fmt.Println(string(byteStr2))
}
