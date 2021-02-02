package main

import "fmt"

// 参数只传递值！
func test(n1 int) { //  n1 是test函数接收 传来值的容器！！！！
	n1 = n1 + 1
	fmt.Println("test(n1)=", n1)
}

func main() {
	num := 20
	test(num) // 此事num= 20 ，函数调用将 20 传给test 然后test用n1 接收！
	fmt.Println("main(num)=", num)
}
