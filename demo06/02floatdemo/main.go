package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//float类型
	var a float32 = 3.12
	fmt.Printf("值：%v--%f, 类型%T\n", a, a, a)

	fmt.Println(unsafe.Sizeof(a)) //float32 占用4个字节

	//float64
	var b float64 = 3.12
	fmt.Printf("值：%v--%f, 类型%T\n", b, b, b)

	fmt.Println(unsafe.Sizeof(b)) //float64 占用8个字节

	// %f 输出float类型  %.2f 输出数据的时候保留2位小数
	var c float64 = 3.1415925535
	fmt.Printf("%v--%f--%.2f\n", c, c, c) // %v原样输出， %f保留6位小数

	//保留4位小数
	fmt.Printf("%.4f\n", c)

	//3、go语言中浮点数默认是float64
	f1 := 3.14156
	fmt.Printf("%f--%T\n", f1, f1)

	//4、Golang科学计数法表示浮点类型
	var f2 = 3.14e2 // 表示f2等于3.14*10的2次方
	fmt.Printf("%v--%f--%T\n", f2, f2, f2)

	var f3 = 3.14e-2 // 表示f3等于3.14除以10的2次方
	fmt.Printf("%v--%f--%T\n", f3, f3, f3)

	//5、float 精度丢失
	var f4 float64 = 1129.6
	fmt.Println(f4 * 100) //输出 112959.99999999999

	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2) //期望是4.4 结果打印出了 4.3999999999999995

	//6、int类型转换成float类型
	x := 10
	y := float64(x)
	fmt.Printf("x的类型是%T, b的类型是%T\n", x, y)

	//7、float类型转换成float类型
	var x1 float32 = 23.4
	x2 := float64(x1)
	fmt.Printf("x1的类型是%T, x2的类型是%T\n", x1, x2)

	//8、float类型转换成int类型
	var c1 float32 = 23.45
	c2 := int(c1)
	fmt.Printf("c2的值：%v, c2的类型：%T\n", c2, c2)
}
