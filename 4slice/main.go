// 1、切片的引入 
//	(1)切片是引用类型
// 	(2)切片使用类似于数组（len、range、）
//	(3)切片长度可以动态变化
// 	(4)切片定义 跟数组一定要分清~~~
//		var 切片名 [] 数据类型
//		举例：var arr []int
package main

import (
	"fmt"
)

func main() {
	// 数组
	var intArr [5]int = [...]int{1, 23, 55, 7, 6}
	// 声明定义一个切片~
	slice := intArr[1:3] // 引用intArr数组中下表从1到3的数字，其中不包含index3
	// [1:3] =[1:3) 意义是 引用 第2个数字到第3个数字！！！！！！！

	fmt.Println("slice=", slice)            //切片的内容
	fmt.Println("slice's len=", len(slice)) //切片的长度
	fmt.Println("slice's cap=", cap(slice)) //切片的容量是动态~~ 自动变大~
	// slice的第一部分就是指向 intArr[1]
	fmt.Printf("slice[0]=%p \n", &slice[0])
	fmt.Printf("intArr[1]=%p \n", &intArr[1])
	fmt.Printf("slice[1]=%v,slice[2]=%v \n", slice[0], slice[1])

	// 2、切片的使用方法
	 // （1）先申明一个数组
	//var intArr1 [5]int = [...]int{1, 23, 55, 7, 6}
	// 然后再引用这个数组的一部产生一个切片~
	//slice1 := intArr1[1:3] // 引用intArr数组中下表从1到3的数字，其中不包含index3

	//  (2) 直接用make来生成切片
	//var 切片名 []数据类型 = make([]type, 长度，容量) 
	// 其中容量cap是可选项目，但是如果你要是写了cap  cap小大于等于len
	var slice01 []int =make([]int, 5)
	slice01[0]=10
	slice01[4]=55
	fmt.Println(slice01)  //10 0  0  0  55
	//	(3) 两个申明方式的区别
	// 		截取数组的切片！数组是可见的，可以直接通过arrP[index]来访问，也可以通过切片来访问
	//		make建立的切片也会产生一个数组，但是是不可见的，只能通过slice[index]来访问，底层维护，程序员不可见。

	// 3、内置函数append   将元素增加到切片的末尾，如果容量足够
	var slice03 []int = []int{100, 23, 55}
	slice03=append(slice03,400,500,600)  // 追加元素
	fmt.Println(slice03)
	slice03=append(slice03,slice01...) 	 // 追加数组，固定写法append（数组，数组名...）
	fmt.Println(slice03) 
// append底层逻辑 熟悉一下
	fmt.Printf("slice03[0]的地址是：%v \n",&slice03[0])
	fmt.Printf("slice03[5]的地址是：%v \n",&slice03[5])
	fmt.Printf("slice03[6]的地址是：%v \n",&slice03[6])

	// 4、copy操作~~ 切片之间才可以，数组不可以
	var a []int=[]int{1,2,3,4}
	var b []int=make([]int,10)
	// 数据是相互独立的~相互不会影响的~
	// 大拷贝到小的里面去，不会报错，有几个位置拷几个
	copy(b,a)
	fmt.Println(a) 
	fmt.Println(b) 
}
