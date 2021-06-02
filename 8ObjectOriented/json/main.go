package main

import (
	"encoding/json"
	"fmt"
)

type Monsters struct {
	Name  string `json:"name"` /*Ttag的重要意义在于让变量以你想要的格式返回*/
	age   int
	skill string
}

func main() {

	monsters1 := Monsters{"IronMan", 30, "robot"}

	//序列化~~跟前端交互的，将monster变量序列化为json格式字符串

	jsonStr, err := json.Marshal(monsters1)

	if err != nil {
		fmt.Println("json处理错误\n", err)

	}

	fmt.Println("jsonmonster", string(jsonStr))

}
