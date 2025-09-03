package main

import "fmt"

func main() {
	//1、基于数组定义切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[:]                        //获取数组里面的所有值
	fmt.Printf("%v  --  %T\n", a, a) // [55 56 57 58 59]
	fmt.Printf("%v  --  %T\n", b, b) // [55 56 57 58 59]

	c := a[1:4]
	fmt.Printf("%v  --  %T\n", c, c) // [56 57 58]

	d := a[2:]
	fmt.Printf("%v  --  %T\n", d, d) // [57 58 59]

	e := a[:3]                       //表示获取第三个下标前面的数据
	fmt.Printf("%v  --  %T\n", e, e) // [55 56 57]

	//2、基于切片定义切片
	a1 := []string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	b1 := a1[1:]
	fmt.Printf("%v  --  %T\n", b1, b1) // [上海 广州 深圳 成都 重庆]

	//3、关于切片的长度和容量
	//容量：切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度%d 容量%d\n", len(s), cap(s)) //长度6 容量6

	a2 := s[2:]
	fmt.Printf("长度%d 容量%d\n", len(a2), cap(a2)) //长度4 容量4

	a3 := s[1:3]
	fmt.Printf("长度%d 容量%d\n", len(a3), cap(a3)) //长度2 容量5

	a4 := s[:3]
	fmt.Printf("长度%d 容量%d\n", len(a4), cap(a4)) //长度3 容量6
}
