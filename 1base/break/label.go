package main

import "fmt"

// break 会跳出最近的for

func main() {
leble1:
	for i := 0; i < 4; i++ {
	leble2:
		for j := 0; j <= 10; j++ {
			if j == 2 {
				break lable2
			} else {
				fmt.Println("j=", j)
			}
		}
	}

}
