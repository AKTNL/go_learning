package main

import "fmt"

func main() {
	//1、请求出一个数组里面元素的和以及这些元素的平均值。分别用for和for-range实现
	var arr = [...]int{12, 23, 45, 67, 2, 5}
	var sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("arr数组元素的和是：%v  平均值：%.2f\n", sum, float64(sum)/float64(len(arr)))

	//2、请求出一个数组的最大值，并得到对应的下标
	var intArr = [...]int{1, -1, 12, 65, 11}
	var max = arr[0]
	var index = 0
	for i := 0; i < len(intArr); i++ {
		if max < intArr[i] {
			max = intArr[i]
			index = i
		}
	}
	fmt.Printf("最大值：%v  最大值对应的索引值：%v\n", max, index)

	//3、从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	var a = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%v,%v)\n", i, j)
			}
		}
	}
}
