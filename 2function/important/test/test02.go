package main

import "fmt"

type mySum func(int, int) int

func sum1(n1, n2 int) int {

	return n1 + n2
}
func sum2(n1, n2, n3 int) int {

	return n1 + n2
}

func myFunc(funcVar mySum, num1 int, num2 int) int {
	return funcVar(num1, num2)

}

func main() {
	a := sum1
	//b := sum2

	fmt.Println(myFunc(a, 1, 2))
	// fmt.Println(myFunc(b, 1, 2)) 是sum2 而sum2 和myFunc 的类型不一致，所以报错

}
