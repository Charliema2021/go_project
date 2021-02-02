package main

import "fmt"

// 创建一个数组，分别放置A-Z。
func main() {
	var letter [26]byte

	for i := 0; i < len(letter); i++ {
		var l byte = 65
		letter[i] = l + byte(i)
		fmt.Printf("%c", letter[i])
	}

}
