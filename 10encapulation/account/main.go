package main

import (
	"fmt"
	"go_project/10encapulation/account/accountInfor"
)

func main() {
	ac1 := accountInfor.NewAccount("gsyh111111", "888888", 200.0) //如果其中一项不满足，就不能继续，跳出一下最好
	// 对账号进行修改操作~~~~
	ac1.SetID("dafdklfdklf")
	fmt.Println(ac1)
}
