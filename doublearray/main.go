package main

import "fmt"

func main() {
	//二维数组声明
	/* 	0 0 0 0 0 0
	   	0 0 1 0 0 0
	   	0 2 0 3 0 0
	   	0 0 0 0 0 0 */

	// arr [4]是列，[6]是行
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	// 二维数组的遍历

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j]) // print 不换行的输出每行元素
		}
		fmt.Println() // 每一行遍历结束后，换行
	}
	// 系统逻辑~~~~
	// 在内存中，arr有两个元素，是两个地址~~指向两个数组
	/* 	arr[0]的地址是：0xc000152000
	arr[1]的地址是：0xc000152030 */
	fmt.Printf("arr[0]的地址是：%p \n", &arr[0])
	fmt.Printf("arr[1]的地址是：%p \n", &arr[1])
	fmt.Printf("arr[0][0]的地址是：%p \n", &arr[0][0])
	fmt.Printf("arr[1][0]的地址是：%p \n", &arr[1][0])

}
