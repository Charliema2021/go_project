package main

import "fmt"

func main() {
	var duoArr [3][5]float64

	for i := 0; i < len(duoArr); i++ { // 多维数组中，len（数组） 等于数组的行
		for j := 0; j < len(duoArr[i]); j++ { // 数组每列的数组 长度，就是列
			fmt.Printf("请输入第%v个班级的第%v个人的成绩:", i+1, j+1)
			fmt.Scanln(&duoArr[i][j])   // 数组的接收 需要逐个接收~~
		}
	}
	// fmt.Println(duoArr)
	// 遍历数组求和~
	   	for i := 0; i < len(duoArr); i++ {
	   		sum := 0.0
	   		for j := 0; j < len(duoArr[i]); j++ {
	   			sum += duoArr[i][j]
	   			//fmt.Printf("arr[%v][%v]=%v", i, j, v2)
	   		}
	   		fmt.Println()
	   		fmt.Printf("第%v个班的总分为%v，平均分%v \n", i+1, sum, sum/float64(len(duoArr[i])))
	   	} 
}
