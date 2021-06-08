package accountInfor

import "fmt"

type accountInfor struct {
	accountID string
	money     float64
	pwd       string
}

// 新建一个结构体~~
func NewAccount(id string, pwd string, money float64) accountInfor {
	var ac accountInfor
	if len(id) < 6 && len(id) > 10 {
		fmt.Println("输入的用户名不符合要求~~~")
	} else {
		ac.accountID = id
	}
	if len(pwd) < 8 {
		fmt.Println("输入的密码不符合要求~~~")
	} else {
		ac.pwd = pwd
	}
	if money < 100 {
		fmt.Println("最少存100元开户~~~")

	} else {
		ac.money = money
	}

	return ac
}

// ac的修改方法~~~
func (ac *accountInfor) SetID(id string) {
	if len(id) < 6 && len(id) > 10 {
		fmt.Println("输入的用户名不符合要求~~~")
	} else {
		ac.accountID = id
		fmt.Println("ID更新完成！\n新ID为：", ac.accountID)
	}

}

func (ac *accountInfor) GetID() string {
	return ac.accountID
}
