package main

import (
	"fmt"
)

// 定义一个结构体，存放学生数据
type studentInfor struct{
// 注意看一下哈~~ 元素名第一个字母大写，才能被其他地方引用哦
	Name string
	Age int
	Sex string
	Address string
}



func Pout(stu map[string]studentInfor){
	for k,v:=range stu{
		fmt.Printf("学号：%v \n",k)
		fmt.Printf("姓名：%v \n",v.Name)
		fmt.Printf("性别：%v \n",v.Age)
		fmt.Printf("籍贯：%v \n",v.Address)
		fmt.Println("--------------------")
	}

}


func main(){
	students := make(map[string]studentInfor,2)	
	stu1:= studentInfor{"matianqi",30,"male","ningxia"}
	stu2:= studentInfor{"zhangsan",25,"female","beijing"}
	students["0001"]=stu1
	students["0002"]=stu2
	
	//fmt.Println(students)
	Pout(students)

}

