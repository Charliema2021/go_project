// 定义一个存放成绩的数组~~~
// 循环输入并最终输出数组相关信息
package main

import (
	"fmt"
	"time"
)

func rangeArray(array [5]float64) {
	//有一个问题，

	for index, value := range array {

		fmt.Printf("[%v] = %v \n", index, value)
		time.Sleep(time.Second)

	}
}

func main() {

	var score [5]float64

	// 循环输入存入数组
	fmt.Println("请输入成绩，输入0结束输入。")
	var over string
	if over != "over" {
		for i := 0; i < len(score); i++ {
			fmt.Printf("请输入第%v个人的成绩=", i+1)
			fmt.Scanln(&score[i])
		}
	} else {
		goto Continue
	}
Continue:
	fmt.Println("--------------------------")
	rangeArray(score)
}
