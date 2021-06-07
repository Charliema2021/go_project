// golang中没有构造模式~所以用工厂模式
// type student struct  这里student实例首字母是小写，但是要再其他包里用这个~~  就用到了工厂模式

 package main

 import (
	"fmt"
	"go_project/9method/factory/student"
 )

func main(){
	// 第一种方法，因为Student 的首字母是大写的，所以再其他包可以直接使用，但是要是小写就会有报错`~~
	// 报错内容如下，意思就是没有这个结构体~~~·
	// 要想正常使用,那就要用到工厂模式
	/* .\main.go:13:11: cannot refer to unexported name student.student
	.\main.go:13:11: undefined: student.student */
/* 	var stu1 student.student
		stu1.Name ="matianqi"
		stu1.Age = 30
		stu1.Score = 95.0 */

	var stu2 = student.NewStu("matianqi~",30,99.0)

	fmt.Println(stu2) // &{matianqi~ 30 98} 用的是地址

	fmt.Printf("name=%v\n",stu2.Name)
	fmt.Printf("age=%v\n",stu2.Age)
	//fmt.Printf("scoer=%v\n",stu2.Score)
	fmt.Printf("score=%v\n",stu2.GetScore()) // 因为小写字母，所以单独给这个score一个方法，让他再student包中传出score
}