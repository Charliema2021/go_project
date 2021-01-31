package main

import (
	"fmt"
)

func Addup() func (int) int{

	var n int =10
	var str string = "Hello"

	return func(x int)int{   //返回值是一个匿名函数，那么n和这个retur构成了一个闭包~，形成一个整体。
		n = n + x
		str += "1"  //自加一下x
		fmt.Println("str=",str)
		return n          // 这个n 会因为调用 一直在改变值
	}
}


func main(){

		f:= Addup()
    // 因为反复调用~初始化一次，然后不断的在改变值，所以才有下面的这个结果。！！！！
		fmt.Println(f(1))   //11
		fmt.Println(f(2))   //13
		fmt.Println(f(3))   //16

	
}