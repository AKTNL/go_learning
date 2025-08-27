package main

import (
	"fmt"
	"strings"
)

func main() {
	//1、定义string类型
	var str string = "你好golang"
	var str2 = "你好go"
	str3 := "你好"

	fmt.Printf("%v--%T\n", str, str)
	fmt.Printf("%v--%T\n", str2, str2)
	fmt.Printf("%v--%T\n", str3, str3)

	//2、字符串转义字符
	s1 := "this \nis str" // 换行
	fmt.Println(s1)

	s2 := "C:\\go\\bin" // C:\go\bin 输出反斜杠\\
	fmt.Println(s2)

	s_2 := "C:Go\"bin" // C:Go"bin
	fmt.Println(s_2)

	//3、多行字符串 （反引号 tab键上面
	st := `this is str
				this is str111
	this is str`
	fmt.Println(st)

	//4、len()求长度
	var sr = "你好"
	fmt.Println(len(sr)) // 长度6

	var sr1 = "aaaa"
	fmt.Println(len(sr1))

	//5、+ 或者 fmt,Sprintf 拼接字符串
	s_str := "你好"
	s_str2 := "golang"
	s_str3 := s_str + s_str2
	fmt.Println(s_str+s_str2, s_str3)

	s_ := fmt.Sprintf("%v %v", s_str, s_str2)
	fmt.Println(s_)

	s11 := "双引号" +
		"照样" +
		"转义字符无效"
	fmt.Println(s11)

	//6、strings.Split 分割字符串  strings需要引入strings包
	var s_1 = "123-456-789"
	arr := strings.Split(s_1, "-")
	fmt.Println(arr) // [123 456 789] 切片，简单的理解切片就是数组 在golang中切片和数组还有一些区别

	//7、strings.join 表示把切片连接成字符串
	str_2 := strings.Join(arr, "*")
	fmt.Println(str_2)

	ar := []string{"php", "java", "golang"} //切片
	fmt.Println(ar)
	str_3 := strings.Join(ar, "-")
	fmt.Println(str_3)
	fmt.Printf("%v - %T\n", str_3, str_3)

	//8、strings.contains 判断是否包含
	sttr := "this is str"
	sttr2 := "this"
	flag := strings.Contains(sttr, sttr2)
	fmt.Println(flag)

	//9、strings.HasPrefix,strings.HasSuffix 前缀/后缀判断
	sttr3 := "this is str"
	sttr4 := "thi"
	sttr5 := "str"
	flag1 := strings.HasPrefix(sttr3, sttr4)
	flag2 := strings.HasSuffix(sttr3, sttr5)
	fmt.Println(flag1, flag2)

	//10、strings.Index(), strings.LastIndex()子串出现的位置 查找不到返回-1,查找到返回下标位置，下标从0开始
	sss := "this is str"
	sss2 := "is"
	num := strings.Index(sss, sss2)
	num2 := strings.LastIndex(sss, sss2)
	fmt.Println(num, num2) // 2 5

}
