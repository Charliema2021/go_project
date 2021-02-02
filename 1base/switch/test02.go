package main

import (
	"fmt"
)

var score uint8

func compare(compare bool) {
	cond0 := fn0(compare)
	if cond0 {
		return
	}

}

func fn0(compare bool) bool {
	if score < 60 {
		return true
	} else {
		return true, compare == false
	}
	return false
}

func main() {
	var score uint8
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)
	switch compare() {
	case true:
		fmt.Println("及格")
	case false:
		fmt.Println("及格啦")
	default:
		fmt.Println("输入错误")
	}

	// 更机智的写法

	switch int(score / 60) {
	case 1:
		fmt.Println("及格了")
	case 0:
		fmt.Println("不格了")
	default:
		fmt.Println("输入有误")
	}

}
