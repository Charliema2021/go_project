package main

import (
	"fmt"
	personinfor "go_project/10encapulation/person/personInfor"
)

func main() {

	per1 := personinfor.NewPerson("matianqi")

	per1.SetAge(18) // set 是往里写入数据
	per1.SetSalary(5000.0)
	//fmt.Println(per1)

	fmt.Println("新建用户信息：")
	fmt.Printf("name=%v ，age=%v ，salary=%v \n", per1.Name, per1.GetAge(), per1.GetSalary())

}
