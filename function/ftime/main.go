package main

import (
	"fmt"
	"time"
)

func main() {
	//time本身也是一个数据类型
	// 1、获取当前时间的方法
	now := time.Now()
	fmt.Printf("now=%v,nowtype=%T \n", now, now)
	// time.now返回了一个结构体：2021-01-30 17:49:51.883501 +0800 CST m=+0.005003001
	// 里面包含了年月日各种信息~~ 直接拿就行了
	// 2、从now中获取年月日时分秒
	fmt.Printf("年=%v \n", now.Year())
	fmt.Printf("月=%v月 \n", int(now.Month())) // 通过int强传，变成数字月份，正常输出是：月=January
	fmt.Printf("日=%v \n", now.Day())
	fmt.Printf("小时=%v \n", now.Hour())
	fmt.Printf("分钟=%v \n", now.Minute())

	// 3、格式化日期的两种方式
	//	（1） 使用Sprintf  字符串输出。就是吧2里面的获取通过格式化输出打出来~~~ 简单了~~
	fmt.Printf("当前时间为%d年%d月%d日，%d点%d分  \n",
		now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute())

	//	（2）fmt   有标准格式！！！！2006年01月02日 15:04:05
	// 	 里面的年月日时间数字不能改动，间隔符可以随意改，就是开发者任性！！！
	fmt.Printf(now.Format("2006年01月02日 15:04:05 \n"))

	fmt.Printf(now.Format("2006-01-02 \n"))

	fmt.Printf(now.Format("15:04:05 \n"))

	fmt.Printf(now.Format("2006"))   //因为这个时间每一个都是不一样的，所以取相关值可以直接写

}
