package main
import (
	"fmt"
	"math/rand"//伪随机
    "time"
)
// 先申明结构体Account
type Account struct {
	AccountNum string
	Pwd string
	Balance float64

}
//创建新用户~

func (a *Account)NewAccount(num string,pwd string){
	//你输入一下哪个银行的，然后我给你随机生成5位数结合你的银行变成
	
	rand.Seed(time.Now().Unix())
	for i:=0;i<=5;i++{
		a.AccountNum = num + string(rand.Intn(10))
	}

	a.Pwd = pwd  // 密码用你给我的密码
	a.Balance = 100.0  // 余额默认初始化都是100

	fmt.Printf("已生成新用户! \n  账户：%v，余额：%v \n",a.AccountNum,a.Balance)

}
// 存钱~~简化实际流程了哦~！~~！~ 以后丰富起来

func (a *Account)SaveMoney(num string,pwd string,ba float64){
	// 先判断存入的钱格式正确
	if ba<=0 {
		fmt.Println("你输入的金额有误！")
		return
	}else{
		fmt.Println("存款成功！")
		a.Balance=a.Balance+ba
		fmt.Println("账户信息如下: \n",a)
	}
	
}

// 取钱~~！~！
func (a *Account)WithdrawMoney(num string,pwd string,ba float64){
	// 先判断取出的钱格式正确
	if ba<=0|| ba>a.Balance {
		fmt.Println("你输入的金额有误！")
		return
	}else{
		fmt.Println("取款成功！")
		a.Balance=a.Balance-ba
		fmt.Println("账户信息如下: \n",a)
	}
	
}




// 查询有个难题~！~！~  怎么遍历 已有的结构体啊~~~
func  (a *Account)SerchAccount(num string,pwd string){
	if pwd ==a.Pwd{
		fmt.Println("账户信息如下: \n",a)
	}else{
		fmt.Println("密码错误！")
	}
} 



func main(){

	var a1= Account{
		AccountNum : "gs111111",
		Pwd : "666666",
		Balance : 800.00,
	}


	// a1.NewAccount("gs","888888")
	
	a1.SerchAccount("gs111111","666666")   // 模拟输入账号密码,也可丰富一下功能哦~~
	a1.SaveMoney("gs111111","666666",100.0)  
	a1.WithdrawMoney("gs111111","666666",200.0)  



}