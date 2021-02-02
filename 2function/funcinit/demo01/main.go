package main

import (
	"fmt"
	"go_project\function\funcinit\demo01\utils"  // 引入包
)

var age = test()

// 为了看到全局变量是被先初始化，
func test() int {

	fmt.Println("test()") //第一个执行
	return 90

}

// init 函数
func init() {

	fmt.Println("init....") //第二个执行
}

func main() {
	fmt.Println("main....age = ", age) //第三个执行
	fmt.Println("Age= ",utils.Age)
	fmt.Println("Name= ",utils.Name)

}
