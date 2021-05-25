// 实现一个银行的基本功能思路
package main

import "fmt"


// 结构体~！~account

type Account struct{
	account string
	passWord string
	balance float32
}


// 存款
func saveMoney(*passWord string,balance float32,add float32){

	// 验证密码
	if passWord != passWord{
		fmt.Println("输入密码有误！")

	}else{
		blance = blance-reduce
		fmt.Printf("账号:%d,余额为:%d",Account.account,balance)
	}

} 

// 取款
func reduceMoney(*passWord string,balance float32,recude float32){

	// 验证密码
	if passWord != passWord{
		fmt.Println("输入密码有误！")

	}else{
		blance = blance-reduce
		fmt.Println("账号余额为:",balance)
	}

} 


// 查询功能

funce checkAccount(*passWord string,balance float32,recude float32){
	// 验证密码
	if passWord != passWord{
		fmt.Println("输入密码有误！")

	}else{
		fmt.Println("账号余额为:",blance)
	}

} 
}




func main(){
// 初始化
	Account.balance=100.0
	Account.passWord="123456"
	Account.account="gsbank1122"

	//查余额
	checkAccount(Account.passWord)



}