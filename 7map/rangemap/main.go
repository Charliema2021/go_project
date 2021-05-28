package main

import (
	"fmt"
)

func main() {
	// For range的方式去遍历
	cities := make(map[string]string)
	cities["1"] = "beijing"
	cities["3"] = "tianjing"
	cities["2"] = "shanghai"

	for key, value := range cities {
		fmt.Printf("key=%v, value=%v \n", key, value)
	}

	// 遍历一个比较复杂的map，多层的时候，多层range

	stuInfor := make(map[string]map[string]string)

	stuInfor["0001"] = make(map[string]string, 4)
	stuInfor["0001"]["name"] = "matianqi"
	stuInfor["0001"]["sex"] = "male"
	stuInfor["0001"]["address"] = "ningxia"
	stuInfor["0001"]["score"] = "A"

	stuInfor["0002"] = make(map[string]string, 4)
	stuInfor["0002"]["name"] = "guojianhong"
	stuInfor["0002"]["sex"] = "famale"
	stuInfor["0002"]["address"] = "beijing"
	stuInfor["0002"]["score"] = "A"

	for k1, v1 := range stuInfor {
		fmt.Printf("number:%v\n", k1)
		for k2, v2 := range v1 {
			//fmt.Printf("name=%v\n sex=%v\n address=%v\n score=%v\n", v2, v2)
			fmt.Printf("%v: %v\n", k2, v2)
		}
		fmt.Println("-----------------")
	}

	// map中到底有多少对key+value，用len
	fmt.Printf("cities 里有%v对内容 \n", len(cities))
	fmt.Printf("stuInfor 里有%v对内容 \n", len(stuInfor))
	fmt.Printf("stuInfor 里有%v对内容 \n", len(stuInfor["0001"]))

}
