// 斐波那契数列
// 传个N给函数，生产一个斐波那契数列，返回一个切片

package main

import "fmt"

func fibona(n int) []uint {

	fbSlice := make([]uint, n)
	//地个数和第二个数都是1
	fbSlice[0] = 1
	fbSlice[1] = 1
	for i := 2; i < n; i++ {
		fbSlice[i] = fbSlice[i-1] + fbSlice[i-2]
	}

	return fbSlice
}

func main() {

	slice := fibona(20)
	fmt.Println(slice)
}
