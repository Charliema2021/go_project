package main

import "fmt"

// 从键盘输入多个数字，计算有多少个正数，多少个负数
func main() {
	var num int
	var Pnum int
	var Nnum int

	for { // 无线循环就用for不加条件
		fmt.Println("请输入数字")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			Pnum++
			continue
		}
		Nnum++
	}
	fmt.Printf("正数的个数为%v，负数的个数为%v", Pnum, Nnum)
}
