package main

import (
	"fmt"
)
//申明一个全局匿名函数
var (
	Func1=func(n1,n2 int)int{
		return n1*n2
	}


)

func main(){
	//定义匿名函数~
	//求两个数的和（使用匿名函数）

	result :=func(n1,n2 int)int{     // 直接定义
		return n1+n2
	}(10,20)                         //这里直接传参

	fmt.Println("result=",result)


	a:=func(n3,n4 int)int{          // 将定义的匿名函数交给一个变量，然后后续用这个变量就行

		return n3-n4
	}
	res2:= a(10,20)         //此时a是一个 函数型的变量~！！~~
	fmt.Println("res2=",res2)


	res3 :=Func1(10,10)
	fmt.Println("result3=",res3)
}