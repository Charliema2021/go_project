// 一个养鸡场有6只鸡，他们的体重分别为，3、5、3.4、2、50。
// 求和、在求平均值
package main

import (
	"fmt"
	"time"
)

// 定义一个遍历数组和切片的函数吧
// 复习一下 数组的遍历
//func rangeArray(array ...interface{})
//这是给balance[] 格式化成为一个字符串传给了另一个函数，输出结果为： [0] = [1000 2 3.4 7 50]
func rangeArray(array []float32) {
	//有一个问题，

	for index, value := range array {

		fmt.Printf("[%v] = %v \n", index, value)
		time.Sleep(time.Second)

	}

}

func main() {

	// 1、定义数组
	//var hens [6] float64;
	// 2、给数组进行初始化
	//	(1)数组初始化的第一种方式，直接在申明时去做 :var variable_name [SIZE] variable_type
	var hens = [6]float64{1.0, 3.0, 5.0, 3.4, 2.0, 50.0}
	//	(2)数组初始化的第二种方式 如果你长度不确定，可以通过[...]来定义，数组会通过输入内容自定义数组长度。
	var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//	(3)指定元素下表的元素的值！！！！
	//	var hens =[3]float32{1:3.2，0:4.0 ，2：5.0}/
	rangeArray(balance)
	// 3、总体重
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	// 4、平均体重
	avgWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))

	fmt.Printf("%v只鸡的总重量为%v，平均体重为：%v", len(hens), totalWeight, avgWeight)

}
