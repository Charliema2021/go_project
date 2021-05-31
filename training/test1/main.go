package main

import "fmt"

//生成随机数
func madearr(arr[]){

	rand.Seed(time.Now().UnixNano()) // 以纳秒级生成随机数
	n := rand.Intn(100) + 1          // 随机生成证书，rand.Intn（范围）


}







func main(){
	var arr [10]int
	madearr(&arr)


}