package main

import "fmt"

func main() {
	a := 9
	b := 5
	// 不引入变量的时候怎么交换a和b的值
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("此时a= %d,b=%d", a, b)

}
