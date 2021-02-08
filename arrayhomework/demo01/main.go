package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1生成一个随机1维数组，10个元素，int类型
// 2倒置数组输出
// 3对其进行标准运算，sum、average、max、min
// 4查找其中有没有55，有就输出位置，没有就输出结果

//  随机数seed生成一个10维数组函数
func createArr(arr *[10]int) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < len(arr); i++ {
		(*arr)[i] = rand.Intn(100)
	}
}

func convert(arr [10]int) {
	var temp int = 0 // 临时变量，用来接收交换用的值
	var len int = len(arr)
	for i := 0; i < len-1; i++ {
		if arr[i] > arr[len-i] { // 退出机制~！！
			temp = arr[i]
			arr[i] = arr[len-i]
			arr[len-i] = temp
		}

	}
	fmt.Println("arr排序后= ", arr)
}

func sumArr(arr [10]int) int {
	var sum int
	for _, value := range arr {

		sum += value

	}
	return sum
}

func main() {
	var arrInt [10]int
	// 生成数组~
	createArr(&arrInt)
	fmt.Println("生成的随机数组为：", arrInt)
	// 数组的转置
	convert(arrInt)

	// jisuan
	sum := sumArr(arrInt)
	fmt.Printf("arrInt的总和为：%v；平均值是：%v \n", sum, sum/len(arrInt))
}
