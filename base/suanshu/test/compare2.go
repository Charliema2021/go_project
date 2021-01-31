package main

import "fmt"

func main() {
	var max int
	var n1 int = 10
	var n2 int = 15
	var n3 int = 20
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	if n3 > max { // 编程思路~~简化步骤，能比少比一次就少一次~！~！~！~
		max = n3
	}
	fmt.Println("max= ", max)
}
