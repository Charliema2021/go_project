package main

import (
	"fmt"
)

//  1匹配后不需要break
func main() {

	// Switch的数据类型要跟case相同
	var n1 int32 = 20
	var n2 int32 = 20 // int 64 会报错
	switch n1 {
	case n2:
		fmt.Println("OK1")
	case 20, 5: // case 后面的表达式可以使多个~用，隔开
		fmt.Println("OK2")
	default: // default不是必须的~
		fmt.Println("没有可匹配的内容")

		// switch 后面不带表达式，可以当成if来用哦~~~
		var age int = 10
		switch {
		case age == 10:
			fmt.Println("OK")
		case age >= 10 && age <= 60: //case 中也可以对范围进行判断
			fmt.Println("适龄哦~")
		}
	}
}
