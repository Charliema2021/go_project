//类型断言
package main

import "fmt"

type Point struct {
	x int
	y int
}
type A interface {
}

func main() {
	var b Point = Point{2, 3}
	var a A // a作为空接口，是可以接受所有内容的～～
	a = b   // b给a之后，a就有了内容，但是他的类型还是interface

	var c Point
	//  c = a  -----所以当这个地方想把a再赋给 c的时候，会报错 类型不一致
	// 解决办法～～～～～ 类型断言
	c = a.(Point) // 意义： 判断a是不是指向point的类型～如果是就*转换*成point并 将值赋给c，若果不是 就报错～～！～！～

	fmt.Println(c)
	// 对类型断言进行一个判断～防止出现panic 中断程序
	var b1 float32 = 3.1
	var a1 A
	a1 = b1                        // 让a1 变成float32～
	if y, ok := a1.(float32); ok { // 在这里申明一个y，让y能够接受a1～～～但是a1是个float32 ～ 对其增加判断～～
		fmt.Println("convert success~")
		fmt.Printf("y的类型是%T，y=[%v]\n", y, y)
	} else {
		fmt.Println("convert fail~")
	}

	fmt.Println("program is working~~~")

}

// 注意事项⚠️
// 1、空接口之前是指向一个类型的，所以后续能够进行类型断言～
