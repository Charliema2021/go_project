package main

import (
	"fmt"
)
// 常用内建函数
// 1、len 获取字符串长度和数组大小，字节！
// 2、new 用来分配内存

func main(){
	// 常规变量申明
	num1:=100
	fmt.Printf("num的类型：%T，num1的值：%v，num1的地址：%v \n",num1,num1,&num1)


	// new 方法
	num2 := new(int)  // *int
	// type 是指针类型，值就是num2的地址，num2地址是存放指针的地址
	fmt.Printf("num2的类型：%T，num2的值：%v，num2的地址：%v，num2指向的值是：%v \n",num2,num2,&num2,*num2)
	// num2 是个指针，他指向了一个地址，这个地址内默认为0

	*num2 =100
	fmt.Println("now num2= ",*num2)

	// 3、make 也可以用来分配内存，主要是给引用类型分配
	// map chanel slice


}
