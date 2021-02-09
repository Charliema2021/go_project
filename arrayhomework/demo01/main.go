package main

import (
	"fmt"
	"math/rand"
	"time"
)





// 1生成一个随机1维数组，10个元素，int类型
//  随机数seed生成一个10维数组函数
func createArr(arr *[10]int) {
	rand.Seed(time.Now().Unix()) // 以当前时间的unix秒作为种子
	for i := 0; i < len(arr); i++ {
		(*arr)[i] = rand.Intn(100) // 随机产生100以内的整数
	}
}
// 2倒置数组输出
// 递归调用一下试试？

func convert(arr [10]int){
	for i:=0;i<len(arr)-1;i++{
		arr[i],arr[len(arr)-1-i]=arr[len(arr)-1-i],arr[i]
	}
	fmt.Println("arr转置后为：",arr)

}


/* func convert(arr [10]int) {
	var temp int = 0 // 临时变量，用来接收交换用的值
	var len int = len(arr)  // 10
	for i := 0; i < len-1; i++ {
		if arr[i] > arr[len-1-i] { // 退出机制~！！
			temp = arr[i]
			arr[i] = arr[len-1-i]
			arr[len-1-i] = temp
		}

	}
	fmt.Println("arr排序后= ", arr)
} */

// 3对其进行标准运算，sum、average、max、min
func sumArr(arr [10]int) int {
	var sum int
	for _, value := range arr {

		sum += value

	}
	return sum
}

// 二分法查找(递归调用)
func serch(arr [10]int, value, left, right, mid int) {

	if left > right {
		fmt.Printf("没找到%v! \n",value)
		return
	}
	mid = (left+right)/ 2
	if value < arr[mid] {
		right = mid - 1
		serch(arr, value, left, right, mid)
	} else if value >arr[mid] {
		left = mid + 1
		serch(arr, value, left, right, mid)
	} else {
		fmt.Printf("找到了%v，位于数组的%v位。 \n", value, mid+1)
	}

}

func main() {
	var arrInt  = [10]int[62,6,13,54,23,52,51,36,37,56]
	// 生成数组~
	//createArr(&arrInt)
	fmt.Println("生成的随机数组为：", arrInt)
	// 数组的转置
	convert(arrInt)

	// jisuan
	sum := sumArr(arrInt)
	fmt.Printf("arrInt的总和为：%v；平均值是：%v \n", sum, sum/len(arrInt))
	// 查找既定数值~ （数组，value，left，right，mid）
	serch(arrInt, 6, 0, len(arrInt)-1, 0)
}
