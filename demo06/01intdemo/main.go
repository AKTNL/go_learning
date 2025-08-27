package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int = 10
	num = 20
	fmt.Printf("num = %v, 类型：%T\n", num, num)

	//int8 范围(-128 ~ 127)
	var n int8 = 98
	fmt.Printf("n = %v, 类型：%T\n", n, n)

	//unsafe.Sizeof 查看不同长度的整型，在内存里的存储空间
	fmt.Println("占用空间：", unsafe.Sizeof(n), "字节") //一个字节

	var a int32 = 15
	fmt.Printf("a = %v, 类型：%T\n", a, a)
	fmt.Println("占用空间：", unsafe.Sizeof(a), "字节")

	//uint类型
	//uint8无符号整型，范围（0～255）
	var b uint8 = 2
	fmt.Printf("b = %v, 类型：%T\n", b, b)
	fmt.Println("占用空间：", unsafe.Sizeof(b), "字节")

	//int类型转换
	var a1 int32 = 10
	var a2 int64 = 21
	fmt.Println(int64(a1) + a2)
	fmt.Println(a1 + int32(a2))

	//高位向低位转换要注意
	var n1 int16 = 130
	fmt.Println(int8(n1)) //-126 有问题

	num1 := 12
	fmt.Printf("num1=%v 类型：%T\n", num1, num1)
	fmt.Println("占用内存空间：", unsafe.Sizeof(num1), "字节") //表示64位的计算机 int就是int64 占用8个字节

	fmt.Printf("num1=%d\n", num1) // %d 表示10进制输出
	fmt.Printf("num1=%b\n", num1) // %b 表示二进制输出
	fmt.Printf("num1=%o\n", num1) // %b 表示八进制输出
	fmt.Printf("num1=%x\n", num1) // %x 表示十六进制输出
}
