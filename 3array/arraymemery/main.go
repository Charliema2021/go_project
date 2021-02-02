package main

import (
	"fmt"
	"time"
)

// array在内存中如何布局哦~~~

func rangeArray(array [3]int64) {
	//有一个问题，

	for index, value := range array {

		fmt.Printf("[%v] = %v \n", index, value)
		time.Sleep(time.Second)

	}
}

func main() {

	var intArr = [5]int64{10, 20, 30, 40, 50}
	//不去初始化的时候，就都是0
	//rangeArray(intArr)
	// int 32每个占用4个字节，所以地址都是 前一个地址+4
	// int64每个占8个字节，所以都是前一个地址+8
	// 数组地址间隔取决于数组的类型决定
	fmt.Printf("intArr[0]的地址%p \n", &intArr[0]) //	xxxxxxx0
	fmt.Printf("intArr[1]的地址%p \n", &intArr[1]) //	xxxxxxx8
	fmt.Printf("intArr[2]的地址%p \n", &intArr[2]) //	xxxxxxx0=xxxxxxx16    ！！！ASCII码里的16对应的是0！！！
	fmt.Printf("intArr[3]的地址%p \n", &intArr[3])
	fmt.Printf("intArr[4]的地址%p \n", &intArr[4])

}
