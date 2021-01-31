package main

import "fmt"

func main() {
	var name string // 用户名
	var pwd string  //
	var loginChance = 3

	for i := 1; i <= 3; i++ {

		fmt.Println("Please input username:")
		fmt.Scanln(&name)
		fmt.Println("Please input password:")
		fmt.Scanln(&pwd)
		if name == "张无忌" || pwd == "888" {
			fmt.Println("登陆成功！")
			break
		} else {
			loginChance--
			fmt.Printf("登陆失败，还有%v次机会输入正确！\n", loginChance)

		}
	}
	if loginChance == 0 {
		fmt.Println("3次输入失败，请核对。")
	}

}
