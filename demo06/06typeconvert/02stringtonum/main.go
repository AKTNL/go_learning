package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		通过fmt.Sprintf()把其他类型转换为String类型
		注意Springf使用中需要注意转换的格式
			int-%d   float-%f  bool-%t   byte-%c
	*/
	var i int = 20
	var f float64 = 12.456
	var t bool = true
	var b byte = 'a'
	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("值：%v--类型：%T\n", str1, str1)

	str2 := fmt.Sprintf("%.2f", f)
	fmt.Printf("值：%v--类型：%T\n", str2, str2)

	str3 := fmt.Sprintf("%t", t)
	fmt.Printf("值：%v--类型：%T\n", str3, str3)

	str4 := fmt.Sprintf("%c", b)
	fmt.Printf("值：%v--类型：%T\n", str4, str4)

	//2、通过strconv 把其他类型转换为string类型

	/*
		FormatInt
		参数1：int64的数值
		参数2：传入int类型的进制
	*/
	str1 = strconv.FormatInt(int64(i), 10)
	fmt.Printf("值：%v   类型：%T\n", str1, str1)

	/*
		参数1：要转换的值
		参数2：格式化类型
			'f'(-ddd.dddd)
			'b'(-ddddp+-ddd,指数为二进制)
			'e'(-d.dddde+=dd,十进制指数)
			'E'(-d.ddddE+-dd,十进制指数)
			'g'(指数很大时用'e'格式，否则'f'格式)
			'G'(指数很大时用'E'格式，否则'f'格式)
		参数3：保留的小数点  -1（不对小数点格式化）
		参数4：格式化类型
	*/
	str2 = strconv.FormatFloat(float64(f), 'f', 3, 32)
	fmt.Printf("值：%v   类型：%T\n", str2, str2)

	str3 = strconv.FormatBool(t)
	fmt.Printf("值：%v   类型：%T\n", str3, str3)

	str4 = strconv.FormatUint(uint64(b), 10)
	fmt.Printf("值：%v   类型：%T\n", str4, str4)

}
