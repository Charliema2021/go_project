package main

import (
	"fmt"
)

func main() {
	i := 1
	for ; i <= 10; i++ { //  for 循环变量初始化；循环条件；循环变量迭代
		//执行顺序，先初始化变量，然后验证条件，然后迭代变量，然后执行循环体，然后验证条件~~~~
		fmt.Println("zhenbbang") //循环体
	}
	// for的第二种写法
	j := 1
	for j <= 10 {
		fmt.Println("dierzhongxiefa", j)
		j++
	}
	// 第三种写法死循环哦~~配合break来使用
	k := 1
	for {
		if k <= 10 {
			fmt.Println("死循环", k)
		} else {
			break
		}
		k++

	}
}
