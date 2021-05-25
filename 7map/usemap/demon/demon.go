package main

import "fmt"

func main() {
	stuInfor := make(map[int]map[string]string)

	stuInfor[0001] = make(map[string]string, 4)
	stuInfor[0001]["name"] = "matianqi"
	stuInfor[0001]["sex"] = "male"
	stuInfor[0001]["address"] = "ningxia"
	stuInfor[0001]["score"] = "A"

	stuInfor[0002] = make(map[string]string, 4)
	stuInfor[0002]["name"] = "guojianhong"
	stuInfor[0002]["sex"] = "famale"
	stuInfor[0002]["address"] = "beijing"
	stuInfor[0002]["score"] = "A"

	fmt.Println(stuInfor)
	fmt.Println(stuInfor[0001])
	fmt.Println(stuInfor[0001]["address"])

}
