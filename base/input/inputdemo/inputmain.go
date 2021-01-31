package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名 ")
	// scanln是等待用户输入，输入后执行下一行
	fmt.Scanln(&name) // 这里取的是name的地址，而不是值，因为现在就没值呀~
	fmt.Println("请输入年龄 ")
	fmt.Scanln(&age)
	fmt.Println("请输入工资 ")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试 ")
	fmt.Scanln(&isPass)
	//if *isPass != true {
	//		return *&isPass =
	//	}
	// 直接用%v 标准格式输出就行了， 哈哈哈不用乱七八糟的了~~！~！
	fmt.Printf("%v的年龄是 %v，收入是 %v，%v通过考试。", *&name, *&age, *&sal, *&isPass)

}
