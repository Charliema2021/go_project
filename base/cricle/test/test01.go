package main

import (
	"fmt"
)

// 打印1~100之间所有是9的倍数的整数及总和
func main() {
	var sum int64 = 0
	var count int64 = 0 // 初始化一个count用来计有多少个数字
	var i int64 = 1
	for i <= 100 { // for循环中第一个变量初始化的那个; 不能省略~！~！~！~！~！
		if i%9 == 0 {
			count++  // 计数函数~~~
			sum += i //自加运算
			fmt.Println(i)
		}
		i++
	}
	fmt.Printf("count=%v,sum=%v", count, sum)
}
