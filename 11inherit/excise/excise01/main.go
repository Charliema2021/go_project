// 匿名字段也可以进行嵌套和继承
package main

import "fmt"

type A struct {
	name string
	Age  int
}

type B struct {
	A
	int // 匿名字段~~
	n   int
}

func main() {
	var b B
	b.int = 20 // 如何使用匿名字段~~~
	b.n = 15   // n是int类型的变量~~ 引用的话也得直接说明

	fmt.Printf("匿名字段int=[%v]\nint变量n=[%v]", b.int, b.n)

}
