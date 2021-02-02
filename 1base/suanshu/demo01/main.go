package main

import "fmt"

func main() {
	// 除法和取余的特色
	//只保留整数部分~

	fmt.Println(10 / 4)
	//
	var n1 float32 = 10 / 4
	fmt.Println(n1)
	// 会根据你赋的值来确定运算后是不是保留小数~~~
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	// 取模~~~~ a%b= a-a/b*b
	fmt.Println("10 % 4 = ", 10%4)     // 10%4= 2余2，取的是余数2
	fmt.Println("-10 % 4 = ", -10%4)   // -10%4= -10-(-10)/3*3
	fmt.Println("10 % -4 = ", 10%-4)   // 10%-4= 10 -10/(-4)*(-4)
	fmt.Println("-10 % -4 = ", -10%-4) // 10%4=

	// ++ --  还是没忘哈哈哈~~·
	// 有两个重点需要注意~！~！
	//		1、只能后++、--没有前
	//		2、不能直接进行比较运算，需要先进行++、--后在进行运算
	var num1 int = 9
	num1++ // 不能用 num1= num1++这种用法。跟JAVA不一样
	//	var i int = 9       ++、-- 只能独立使用 你先++ 在赋值给他
	//	var a int      		i++
	//	a= i++	 			a = i
	var a int = num1
	fmt.Println("num1++=", num1)
	fmt.Println("a=", a)

}
