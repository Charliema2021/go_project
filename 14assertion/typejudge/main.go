package main

import "fmt"

// 预制一个空接口来接受内容～
type itsm interface {
}

type Student struct {
}

func TypeJudge(itsm ...interface{}) {
	for index, value := range itsm {
		switch value.(type) {
		case int:
			fmt.Printf("第%v个内容的类型是%T，值为%v\n", index+1, value, value)
		case float32:
			fmt.Printf("第%v个内容的类型是%T，值为%v\n", index+1, value, value)
		case float64:
			fmt.Printf("第%v个内容的类型是%T，值为%v\n", index+1, value, value)
		case bool:
			fmt.Printf("第%v个内容的类型是%T，值为%v\n", index+1, value, value)
		case string:
			fmt.Printf("第%v个内容的类型是%T，值为%v\n", index+1, value, value)
		case Student: // 也可以判断自定义类型的
			fmt.Printf("第%v个内容的类型是%T，值为%v\n", index+1, value, value)
		case *Student:
			fmt.Printf("第%v个内容的类型是%T，值为%v\n", index+1, value, value)
		default:
			fmt.Println("无法识别的类型～～")
		}
	}

}

func main() {
	var its [7]itsm
	its[0] = 100
	its[1] = "matianqi"
	its[2] = 98.5
	its[3] = true
	its[4] = "hello"
	var stu1 = Student{}
	var stu2 = &Student{}
	its[5] = stu1
	its[6] = stu2

	for _, val := range its {
		TypeJudge(val)
	}
}
