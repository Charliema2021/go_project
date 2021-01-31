package main

import "fmt"

func main() {
	//
	var i int = 10
	//i的实际地址怎么输出（门牌号）&
	//每次输出的时候存的都不是一个地址哦~
	fmt.Println("Adress is ", &i)

	// ptr里面存的是&i 的地址，通过&i的地址取到 10
	var ptr *int = &i
	// 指针变量只能存地址，不能存具体值
	// 指针类型要匹配 比如 int 和float 不能互相匹配
	fmt.Println("ptr= ", ptr)
	fmt.Println("ptr Adress", &ptr)

	// 获取指针 指向的空间的值  *
	fmt.Printf("ptr指向的值=%v", *ptr)

}
