package main

import "fmt"

func bubbleSort(arr *[5]int) {
	var temp int = 0 // 临时变量，用来接收交换用的值
	var len int = len(*arr)
	for j := 1; j <= len-1; j++ {
		for i := 0; i < len-j; i++ {
			//第一轮排序交换
			/* 		&[24 69 80 57 13]
			   		&[24 69 80 57 13]
			   		&[24 69 57 80 13]
			   		&[24 69 57 13 80] */
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}

		}
	}
	fmt.Println("arr排序后= ", (*arr))
}

func main() {

	arr := [5]int{24, 69, 80, 57, 13}
	bubbleSort(&arr) // 传地址去那边哦，直接影响原来的函数
}
