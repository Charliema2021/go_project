package main

import (
	"fmt"
	"sort"
)

func main() {
	var num map[int]int = make(map[int]int, 10)
	num[1] = 10
	num[2] = 82
	num[30] = 12
	num[14] = 14
	num[25] = 13
	num[6] = 196
	fmt.Println(num)

	// map 的排序思路
	// 先将map的key放入到一个切片中

	var keyOrder []int
	for k, _ := range num {
		keyOrder = append(keyOrder, k)
	}

	// 对切片进行排序,这里用的是一个内置函数，由低到高s排序
	sort.Ints(keyOrder)
	fmt.Println(keyOrder)

	// 遍历切片~按照key输出map的纸值

	for _, v := range keyOrder {
		fmt.Printf("num[%v]=%v \n", v, num[v])
	}

}
