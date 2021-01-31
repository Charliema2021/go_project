package main

import (
	"fmt"
)

func main() {
	var num int = 10

	switch num {
	case 10:
		fmt.Println("Ok1")
		fallthrough // 默认只能穿透一层，一般用不着哦
	case 28:
		fmt.Println("OK2")

	}

}
