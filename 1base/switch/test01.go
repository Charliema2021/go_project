package main

import (
	"fmt"
)

func main() {
	var letter string // 存成byte 有什么特殊的
	fmt.Println("请输入小写字母：")
	fmt.Scanln(&letter) // scanf（“%c”，&letter）

	switch letter {
	case "a":
		fmt.Println("A")
	case "b":
		fmt.Println("B")
	case "c":
		fmt.Println("C")
	default:
		fmt.Println("不匹配")

	}

}
