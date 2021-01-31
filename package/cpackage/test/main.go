// 写一个程序
// 调用闭包，函数判断某个字符串是否有指定的后缀，如果有输出原名，如果没有，加上jpg
package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string)string{
//   此时闭包函数为：
//		变量是suffix
//		匿名函数是makeSuffix（suffix） 
	return func (name string)string{
		// 如果没有指定的后缀名，就加上
		if !strings.HasSuffix(name,suffix) { //如果 传入给name的名字后缀不一样。
			
			return name + suffix   //加上这个.jpg这个后缀
		}else{

			return name
		}
	}
}


func main(){

	m:= makeSuffix(".jpg")    //这个函数需要你传进去一个指定的后缀。方式如此~
// 使用闭包的好处是，验证这种东西，我只需要把核心变量传一次就行，保留住，不用反复传。
	fmt.Println("文件名处理后结果为：",m("123.vai"))
	fmt.Println("文件名处理后结果为：",m("123~"))
	fmt.Println("文件名处理后结果为：",m("123.jpg"))



}