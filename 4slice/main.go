// 1、切片是引用类型
// 2、切片使用类似于数组（len、range、）
// 3、切片长度可以动态变化
// 4、切片定义 跟数组一定要分清~~~
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
	fmt.Printf("slice[1]=%v,slice[2]=%v", slice[1], slice[2])
}
