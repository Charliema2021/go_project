// *5、编写一段代买来统计函数执行的时间
package main

import (
	"fmt"
	"time"
	"strconv"
)

func demo01(){

	str:=""
	for i:=0;i<10000;i++ {
		str +="hello" + strconv.Itoa(i) // 将i强转成string 进行拼接循环拼接
	}

}

func main(){
	
	start:= time.Now().UnixNano()
	demo01()
	end:= time.Now().UnixNano()
	fmt.Println("耗费时间为：",end-start)
// 太快了 所以秒级无法显示，用了纳秒  55972300


}