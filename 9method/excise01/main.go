package main

import "fmt"

//编写一个结构体，给他一个发发，打印一个10*8的矩形
type MethodUitls struct {
}

func (m MethodUitls) print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Println("*")
		}
		fmt.Println()
	}

}

func main() {
	var m MethodUitls
	m.print()

}
