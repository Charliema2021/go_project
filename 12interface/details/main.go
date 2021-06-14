package main

import "fmt"

type A interface {
	Say()
}

type B interface {
	Teach()
}

type Student struct {
	Name string
}

func (stu Student) Say() {
	fmt.Println("student is saying~~~")
}

// 不只是结构体~~常规变量也可以使用接口~只要她实现了接口定义的方法就行
type integer int

func (t integer) Say() {
	fmt.Println("integer is saying~~~~~~", t)
}

type teacher struct {
}

func (t teacher) Teach() {
	fmt.Println("正在上课中~~~")
}
func (t teacher) Say() {
	fmt.Println("老师正在说话中~~~")
}
func main() {
	var stu Student
	A.Say(stu)
	// 这里提供了两种使用接口的方式~~
	//1、var 接口变量，然后把实现这个方法的变量给a~~
	var a A = stu //只要这个类型实现了say就能给传给接口变量
	a.Say()
	// 直接var i，然后调用接口方法~~把A传过去~
	var i integer = 10
	A.Say(i)

	// 变量和接口之间的关系是！ 变量必须完成一个接口中的所有方法才能调用，反之一个变量可以完成多个接口~~~
	var teacher01 teacher
	A.Say(teacher01)
	B.Teach(teacher01)
}
