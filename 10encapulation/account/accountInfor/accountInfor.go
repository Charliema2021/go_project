package accountInfor

import "fmt"

type accountInfor struct {
	accountID string
	money     float64
	pwd       string
}

// 新建一个结构体~~
func NewAccount(id string, money float64, pwd string) {
	if len(id) < 6 && len(id) > 10 {
		fmt.Println("输入的用户名不符合要求~~~")
	}

}

func (ac *accountInfor) SetID(id string) {
	if len(id) < 6 && len(id) > 10 {

	} else {
		ac.accountID = id
	}

}

func (ac *accountInfor) GetID() string {
	return ac.accountID
}
