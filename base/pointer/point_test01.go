package main

import "fmt"

func test01() {
	var num int = 9
	//fmt.Printf("num= ", num)
	fmt.Printf("num adress =%v", &num)
	var ptr *int //申明一个指针变量
	ptr = &num   //给ptr一个num的地址
	*ptr = 10    // 找到ptr对应的存储位置，赋一个10
	fmt.Printf("num=%v", num)
}
