package main

import "fmt"

// 赋值运算符~ =
// 复核赋值运算符~ +=、-=、*=、/=、%=   先运算再赋值 【例】 A *= C  把A*C 赋给A

func main() {
	var i int = 10
	var j int = 9
	// 两个变量交换赋值~
	t := i // 这里就不用在var了，因为t：= 已经等于声明了
	i = j
	j = t
	fmt.Printf("i=%d,j=%d", i, j)
}
