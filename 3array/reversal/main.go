//  随机得到一组数 存到一个数组中
//  反转打印该数组
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 交换元素
func change(arr []int) {
	fmt.Println(arr)
	var temp int = 0
	len := len(arr)
	for i := 0; i < len; i++ {
		temp = arr[i]         // 把第一个数字给temp！
		arr[i] = arr[len-1-i] // 把最后一个数字 给第一个数字，倒数第二个，给正数第二个！
		arr[len-1-i] = temp   // 把temp里的数字给第一个！
	}
	fmt.Printf("arr type=%T",arr)

}

// 输出数组
func readArr(arr [5]int) {
	for index, value := range arr {
		fmt.Printf("[%v]=%v \n", index, value)
	}
	// 忽略一些取值可以直接用  “_”
	/* 	for _, value := range arr {
		fmt.Printf("%v \n", value)
	} */
}

//

func main() {
	// 1、生成随机数组
	var intArr [5]int
	Len := len(intArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Len; i++ {
		// seed 放在循环体外面，不然放在循环体里面会导致每次循环进来都是同一个seed
		// 随机数生成逻辑是，先有一个seed然后以此为依据生成第二个数字，第二个数字再生成第三个数字
		/*
		   rand.Seed:
		       还函数是用来创建随机数的种子,如果不执行该步骤创建的随机数是一样的，因为默认Go会使用一个固定常量值来作为随机种子。

		   time.Now().UnixNano():
		       当前操作系统时间的毫秒值
		*/
		intArr[i] = rand.Intn(100) // 在没有提供随机数种子之前，生成的随机数永远是[81 87 47 59 81]
	}

	// 2、输出原数组
	//readArr(intArr)
	fmt.Println("数组转置后")
	// 3、交换数组
	slice :=intArr[:]
	change(slice)
	// 3、输出数组

}
