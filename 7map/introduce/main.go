package main

import "fmt"

func main() {
	// map 申明
	var a map[string]string
	a = make(map[string]string, 10)
	a["n.1"] = "AAA"
	a["n.2"] = "BBB"
	a["n.1"] = "CCC" // key是不可以重复的~ 如果重复的key会导致覆盖
	a["n.3"] = "BBB" // value是可以重复的
	//在不make初始化时候，直接用make是没法使用的，报错如下
	//panic: assignment to entry in nil map
	fmt.Println(a)
}

// 再golang里map是一个无序的~！！~！~其他语言中会有排序哦！！ 有序啊~~？
