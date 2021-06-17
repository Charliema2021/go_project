package main

import (
	"fmt"
	"sort"
)

func main() {
	//定义一个切片
	var intSlice []int =[]int{10,-1,50,5,56}

	// 对切片进行排序
	// 1、冒泡排序
	// 2、用sort 排序

	sort.Ints(intSlice)
	fmt.Println(intSlice)
}