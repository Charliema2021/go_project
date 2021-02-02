package main

import (
	"fmt"
)

// 按照格式输出矩阵

func main() {
	var val int64
	fmt.Println("请输入数字：")
	fmt.Scanln(&val)
	var i int64 = 0
	for ; i <= val; i++ {
		fmt.Printf("%v + %v = %v \n", i, val-i, val)
	}

}
