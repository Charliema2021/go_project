package main

import "fmt"

// 定义一个结构体，然后给这个结构体一个方法计算面积并返回

type Circle struct {
	Radius float64
}

func (c *Circle) area() float64 { // c在这里是一个接收main传来参数的元素，自己的元素自己用~~~~
	// 这里是值传递，所以这里运算用的c 是r的考呗，如果变成指针，就变成地址传递，这个时候r是会再area中被修改的

	return (*c).Radius * (*c).Radius * 3.14

	// 为了提高效率，通常方法和结构体指针类型绑定！！！！
	// 实际操作的时候会有部分省略，但是真实传递还是指针类型
	//func (c *Circle) area() float64{
	//	return c.Radius * c.Radius * 3.14  这里的c.Radius还是(*c).Radius,只不过golang有简写
	//}
}

func main() {
	var r Circle
	r.Radius = 3.0
	S := (&r).area()

	fmt.Printf("r=%v ，S= %v \n", r, S)

}
