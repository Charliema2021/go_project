//	结构体引入
package main

import "fmt"

type cat struct {
	// 注意看一下哈~~ 元素名第一个字母大写，才能被其他地方引用哦
	Name   string
	Age    int
	Colour string
}

func main() {
	//创建cat的数据啦
	var cat1 cat
	cat1.Name = "小白"
	cat1.Age = 2
	cat1.Colour = "白色"

	fmt.Println(cat1)
	fmt.Printf("cat1的地址是%p \n", &cat1)
}
