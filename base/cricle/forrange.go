//字符串遍历方式1
package main

import (
	"fmt"
)

func main() {
	//var str string = "hello,world!北京你好~"
	// 传统字符串遍历是一个字节一个字节遍历的，所以如果是汉字（占3个字节），就无法正常取出
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c\n", str[i])
	//}
	// 引入切片 rune[]
	var str string = "hello,world!北京你好~" // string 类型是utf-8，一个字节~
	str2 := []rune(str)                  // 强制转换，将str转换成切片rune（一种数据类型） 然后赋值给str2~！~！~！
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	str = "abcdefg大哥你好"
	for index, val := range str {
		fmt.Printf("index=%v,val=%c \n", index, val) // 汉字占3个字节，所以index输出会展现出来
	}
}
