package main

import "fmt"

func main() {
	// 第一种方式
	var a map[string]string
	a = make(map[string]string, 10)
	a["n.3"] = "CCC"
	a["n.2"] = "BBB"
	a["n.1"] = "AAA"
	fmt.Println(a)

	// 第二种方式（推荐使用）:结构清晰
	cities := make(map[string]string)
	cities["1"] = "beijing"
	cities["3"] = "tianjing"
	cities["2"] = "shanghai"
	fmt.Println(cities)

	// 第三种方式
	// 简写heroes := map[string]string{}
	var heroes map[string]string = map[string]string{
		"heroes1": "Iron Man",
		"heroes2": "Captain America", // , 不能丢掉
	}
	// 后续需要拓展的时候,配合使用
	heroes["3"] = "spider man"
	fmt.Println(heroes)

}
