package main

import (
	"fmt"
)
// 对数据进行处理

func modifyUsers(users map[string]map[string]string,name string){
	users := make(map[string]map[string]string, 1)
	fmt.Println(users)
	//fmt.Println(name)
	//判断现有的map中有没有目标元素
/* 	v,ok:=users[name]
	if ok{
		// 有这个用户,对密码进行修改
		users[name]["pwd"]="888888"


	}else{
		// 没有这个用户，进行新建
		users[name]=make(map[string]string,2)
		users[name]=name
		users[name]["pwd"]="66666"
		users[name]["nickName"]=name+"~昵称"

	} */


}


func main(){

	users := make(map[string]map[string]string, 2)
	
	users["matianqi"]["pwd"]="111111"
	users["matianqi"]["nickName"]="~nichengma"

	//fmt.Println(users)

	

	modifyUsers(users,"matianqi")
	//modifyUsers(users,"zhangsan")

	//fmt.Println(users)



}