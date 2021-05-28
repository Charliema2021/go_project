// var map []map
// 意义在于动态的map 可以随意增加
// 用map记录一组monsters~~~ 并且妖怪的个数可以变化
package main

import (
	"fmt"
)

func main() {
	// 申明一个map切片
	var monsters []map[string]string = make([]map[string]string, 2)

	// 切片初始化空间也得make一下，别忘了
	// nil~~~~切片都是 monsters[index]=value
	monsters[0] = make(map[string]string, 2) // 这个时候monsters[0]是一个单独的map，只要是map就得先make一下
	monsters[0]["name"] = "niumowang"
	monsters[0]["age"] = "500"

	monsters[1] = make(map[string]string, 2)
	monsters[1]["name"] = "zhizhu"
	monsters[1]["age"] = "800"

	/* 	if monsters[0] == nil {
	   		monsters[0] = make(map[string]string, 2)
	   		monsters[0]["name"] = "niumowang"
	   		monsters[0]["age"] = "500"
	   	}
	   	if monsters[1] == nil {
	   		monsters[1] = make(map[string]string, 2)
	   		monsters[1]["name"] = "zhizhu"
	   		monsters[1]["age"] = "800"

	   	} */

	//切片有自己的内置函数，可以实现增加的功能
	// append
	var addMonster map[string]string = make(map[string]string)
	addMonster["name"] = "shenggongbao"
	addMonster["age"] = "200"

	monsters = append(monsters, addMonster)
	fmt.Println(monsters)

}
