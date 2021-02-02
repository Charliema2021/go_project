package main

import "fmt"

func test01(arr [3]int) {

	arr[0] = 88
	fmt.Println("arr`1=", arr)

}
func test02(arr *[3]int) { // test-02 用指针数组接手这组数组的地址

	(*arr)[0] = 88 // 拿到数组第一个的地址，进去给他重新赋值88   拿地址的方式（*arr）
	fmt.Println("arr`2=", arr)

}

func main() {

	// 1、数组一旦声明了，长度固定，同时内涵元素要相同
	var arr01 [3]int

	arr01[0] = 1
	arr01[1] = 1
	//arr01[2] = 1.2  // 1.2 (untyped float constant) truncated to int
	arr01[2] = 3
	fmt.Println(arr01)
	// 2、切片和数组要分开哦~ var arr []int  这是个切片

	// 3、数组申明后，不初始化会有基础赋值
	/* 	（1）int类型 初始值是0
	（2）string类型 初始值是 " "
	（3）bool类型 初始值是 false
	*/
	var arr3_1 [3]int
	var arr3_2 [3]string
	var arr3_3 [3]bool

	fmt.Println(arr3_1)
	fmt.Println(arr3_2)
	fmt.Println(arr3_3)
	// 4、数组的下标是从0开始的，如果超过得报越界的错误~~~ panic

	// 5、go的数组属于值拷贝，所以传过去的是copy过去的，不会影响到原数组。！！
	arr05 := [3]int{1, 2, 3}
	test01(arr05)
	test02(&arr05) // 把地址传给test02
	fmt.Println("arr05=", arr05)

	// 6、传递数组的时候要注意，数组长度也是其中一部分，也要传过去哦~
	// 去看一下extreme 里的函数调用，切片和数组要分开~！

}
