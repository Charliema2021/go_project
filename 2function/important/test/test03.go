//在函数中交换两个变量的值
package main

import "fmt"

func swap(n1, n2 *int) {

	t := *n1
	*n1 = *n2
	*n2 = t
	// 回忆一下之前，不引入新的变量怎么换两个函数的值
}

func main() {

	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a=%v,b=%v", a, b)

}
