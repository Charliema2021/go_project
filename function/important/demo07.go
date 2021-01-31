package main

import "fmt"

func sum(n1 int, args ...int) (sum int) {
	sum = n1
	//对args遍历
	// args这个参数名可以自定义。约定俗成用args
	for i := 0; i < args[i]; i++ {

		sum += args[i]
	}
	return sum
}

func main() {

	res := sum(10, 20, 30, 50)
	fmt.Println("result=", res)

}
