package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) test() { // 通过定义方法给Person 定一个专用的方法

	p.Name = "jack" // 方法也是值传递~~~~！！！！ 所以对于他的操作无法改变其的值。
	fmt.Println("test().Name=", p.Name)

}

// 传参试试
func (p Person) jisuan(num int) {
	var result int
	for i := 0; i <= num; i++ {
		result = i
	}
	fmt.Println(result)

}

//传参，再返回值
func (p Person) getSum(a, b int) int {
	return a + b
}

func main() {

	var per1 Person
	per1.Name = "tom"
	per1.Age = 10

	per1.test() // 专有的调用方法
	fmt.Println(per1)

	per1.jisuan(10)
	res := per1.getSum(7, 8)
	fmt.Println("sum=", res)
}
