// 提高一下难度~~~
// 数组随机生成，然后求和
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(arr [5]int) int {
	var Sum int
	for _, value := range arr {

		Sum += value

	}
	return Sum
}

func randArr(arr *[5]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(*arr); i++ {
		arr[i] = rand.Intn(100)
	}
}

func main() {

	var arrayMain [5]int
	randArr(&arrayMain)

	fmt.Printf("随机数组为：%v \nsum=%v", arrayMain, sum(arrayMain))

}
