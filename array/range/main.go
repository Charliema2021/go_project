package main

import (
	"fmt"
)

func main() {

	var hens = [6]float64{1.0, 3.0, 5.0, 3.4, 2.0, 50.0}

	fmt.Println(hens)
	// range返回的值是两个，需要用两个变量接收
	for index, value := range hens {
		fmt.Printf("[%v]=%v \n", index, value)
	}
	// 忽略一些取值可以直接用  “_”
	for _, value := range hens {
		fmt.Printf("%v \n", value)
	}

}
