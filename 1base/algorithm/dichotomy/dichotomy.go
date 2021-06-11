package main

import "fmt"

func main() {

	strList := "abcdefghigk"
	low := 0
	high := len(strList)
	gusee := "h"
	if low <= high {
		mid := (low + high) / 2
		if gusee == &[mid] {
			fmt.Println("找到了h")

		}

	}

}
