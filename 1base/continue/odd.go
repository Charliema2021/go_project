package main

import "fmt"

// 打印所有的偶数~
func main() {

	for i := 1; i <= 100; i++ {
		if i%2 == 0 { // i%2==0 表示能被2整除的数字
			fmt.Println("i=", i)
			continue // 继续执行for循环，只跳出当前if
		}
	}
}
