package main

import "fmt"

func sum(n1, n2 float32) float32 {
	fmt.Printf("n1 type=%T \n", n1) // 浮点数也能计算的~ 懵逼了
	return n1 + n2
}

func main() {
	fmt.Println("sum=", sum(1, 2))
}
