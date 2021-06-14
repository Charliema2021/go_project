package main

import "fmt"

type A interface {
	Yuwen()
}

type B interface {
	English()
}

type nil interface { // 空接口可以被任何类型都实现，所以可以把任意变量都赋给他~

}
type exams interface { // 要想继承A和B的接口~ 那么就需要全部实现A和B中的所有方法~~~
	A
	B
	score()
}

type student struct {
}

func (stu student) Yuwen() {
	fmt.Println("参加语文考试中~~~~")
}
func (stu student) English() {
	fmt.Println("参加英语考试中~~~~")
}
func (stu student) score() {
	fmt.Println("得分为~~~~")
}
func main() {

	var stu student
	var a A = stu        // 证实A被完全实现
	var b B = stu        // 证实B被完全实现
	var exam exams = stu // 证实exam中完全继承了A和B 被完全实现
	a.Yuwen()
	b.English()
	exam.Yuwen()
	exam.English()
	exam.score()
	// 如何使用空接口~

	var num1 float64 = 10.01
	var N nil = num1
	fmt.Println(N)
}
