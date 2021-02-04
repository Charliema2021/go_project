package main

import "fmt"

func main() {
	// 1、string 也可以通过切片进行引用~~~

	var str string = "nihaoya~"
	// 切片引用规则还是 那样的左闭右开
	slice := str[0:3]
	fmt.Println(slice)

	// 2、string是不可变的，本身不能用下标记进行赋值
	// str[0] = "z"    cannot assign to str[0]

	// 3、那你想要改变string，步骤是先转成切片，然后在变成string

	/* 	slice01 := []byte(str)
	   	slice01[0] = 'W' // 修改也是等量改的 1对1 ，byte
	   	str = string(slice01)
	   	fmt.Println(str) */
	slice02 := []rune(str)
	slice02[0] = '北' // 如果想改变汉字，就处理成rune
	fmt.Println(str)
}
