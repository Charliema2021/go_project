package main

import "fmt"

// 一般不用goto因为容易造成程序循环混乱
func main() {
	fmt.Println("试试goto1")
	goto label1
	fmt.Println("试试goto2")

label1:
	fmt.Println("试试goto3")
	fmt.Println("试试goto4")

}
