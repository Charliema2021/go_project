package main

import (
	"fmt"
)

func test4(n1 int, n2 int) int {
	return n1 + n2
}

func test41(funvar func(int, int) int, num1 int, num2 int) int { // 形参funvar 申明为函数类型，你来时候的test4是一个func（num1，num2）所以我func也得带俩跟形参
	return funvar(num1, num2)
}

func main() {

	a := test4 //a 用来接收test4的数据类型的~
	fmt.Printf("sum的数据类型%T，test4数据类型为%T  \n", a, test4)

	sum := test4(1, 2) // 调用test4 传入1和2
	fmt.Printf("sum=%v   \n", sum)

	sum2 := test41(test4, 50, 60)
	fmt.Println("sum2=", sum2)
}
