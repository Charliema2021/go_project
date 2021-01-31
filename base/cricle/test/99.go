package main

import "fmt"

func main() {
	var len int // 总层数
	fmt.Println("请输入需要的层数：")
	fmt.Scanln(&len)

	for i := 1; i <= len; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v=%v \t", j, i, j*i)
		}
		fmt.Println() //换行
	}

}
