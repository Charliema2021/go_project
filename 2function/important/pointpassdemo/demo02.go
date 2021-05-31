package main

import "fmt"

func test(n1 *int) { // 此时n1申明为* 指针类型
	*n1 = *n1 + 1
	fmt.Println("test(n1)=", *n1)
	fmt.Printf("n1的地址是%v \n", &n1)
}

func main() {
	num := 20
	test(&num) // 这次用地址符 将main中&num= 20的地址传给test
	fmt.Println("main(num)=", num)
	fmt.Printf("num的地址是%v", &num)
}
