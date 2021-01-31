package main

import (
	"fmt"
)

func main() {
	type myInt int // 给int换个名字叫myInt
	var num myInt  // 此时num的类型是main.myInt
	num = 20
	//然而在实际中，电脑并不认为int和myInt并不是同一个类型！
	var num2 int
	num2 = int(num) // vs自动强转了~~ 将num强转为int

	fmt.Printf("num的类型是啥 %T \n", num)
	fmt.Printf("num2的类型是啥 %T", num2)
}
