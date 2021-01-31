package main

import (
	"fmt"
)

//可以通过在申明函数时，直接定义返回值，这样就直接return ，按照顺序返回哦~~
func getSum(n1 int, n2 int) (sum int, sub int) { //   先是sum 然后sub
	// 申明时候定义了顺序，这样在函数内，先算谁，后算谁都无所谓。
	sub = n1 - n2
	sum = n1 + n2
	return // 直接return就行  不用管返回谁了

}
func main() {

	a, b := getSum(20, 10) // a和b 是顺序接收。先sum 后sub
	//a,_ :=  getSum(20, 10)   "_" 是一个占位符，你如果只需要接一个特定值就用这个占位符

	fmt.Println("sum = ", a)
	fmt.Println("sub = ", b)

}
