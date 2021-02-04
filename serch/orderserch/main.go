package main

import "fmt"

func main() {
	name := [5]string{"白眉鹰王", "青翼蝠王", "紫衫龙王", "金毛狮王"}
	var heroName string = "" // 定义一个 接收查询内容的地方
	fmt.Println("请输入需要查询的人名字：", heroName)
	fmt.Scanln(&heroName)

	for i := 0; i < len(name); i++ {
		if heroName == name[i] {
			index := i // 已经i++了，所以这个时候的i已经是真实顺序了
			fmt.Printf("找到了%v在数组的第%v个！", heroName, index+1)
			break
		} else if i == (len(name) - 1) {
			fmt.Printf("没有找到%v", heroName)
		}
	}
}
