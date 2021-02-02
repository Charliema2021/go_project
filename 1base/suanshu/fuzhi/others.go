package main

import "fmt"

func main() {

	a := 100
	fmt.Println("a的地址是", &a)

	var ptr *int = &a
	fmt.Println("ptr指向的值是", *ptr)
	fmt.Println("ptr指向的地址是", &ptr)
}
