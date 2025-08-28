package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1、string类型转换成整型
	str := "123456"
	fmt.Printf("%v--%T\n", str, str)
	/*
		ParseInt
		参数1：string数据
		参数2：进制
		参数3：位数 32 64 16
	*/
	num, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%v--%T\n", num, num)

	/*
		ParseFloat
		参数1：string数据
		参数2：位数 32 64
	*/

	str = "123456.333"
	num1, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("%v--%T\n", num1, num1)

	//不建议把string类型转换成bool型
	b, _ := strconv.ParseBool("true")
	fmt.Printf("%v--%T\n", b, b)
}
