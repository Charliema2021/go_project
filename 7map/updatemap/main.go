package main

import "fmt"

func main() {
	// map ["key"] = value
	// 这个过程中，如果key存在，则更改value，没有key，则就是增加
	cities := make(map[string]string)
	cities["1"] = "beijing"
	cities["3"] = "tianjing"
	cities["2"] = "shanghai"
	fmt.Println(cities)

	cities["3"] = "~tianjing~"
	fmt.Println(cities)

	// *map的删除
	//1、有key的时候，进行删除操作
	// 用的是delete这个内置函数进行的~！~！~！
	delete(cities, "1")
	fmt.Println(cities)
	//2、没有这个key的时候不报错，不进行操作
	/* 	delete(cities, "4")
	   	fmt.Println(cities) */

	// 3、golang中没有一次性清空所有的key,只能通过遍历删除来
	//	或者重新make一次 这个map，让之前的数据变成辣鸡，被pc回收
	/* 	cities = make(map[string]string)
	   	fmt.Println(cities) */

	// map的查找
	val, ok := cities["2"]
	if ok {
		fmt.Printf("the key is exist,the value is %v \n", val)

	} else {
		fmt.Printf("the key is not exist \n")
	}
}
