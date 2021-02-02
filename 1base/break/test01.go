package main

import "fmt"

//  100以内的数求和，当和第一次大于20的当前数
func main() {

	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println(i)
			break
		}
	}

}
