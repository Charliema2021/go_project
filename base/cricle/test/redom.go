package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 0

	for {

		rand.Seed(time.Now().UnixNano()) // 以纳秒级生成随机数
		n := rand.Intn(100) + 1          // 随机生成证书，rand.Intn（范围）
		fmt.Println("n=", n)
		count++

		if n == 99 {
			break // 当生成树=99时，调出For循环~~~
		}

	}
	fmt.Printf("生成99以供用了 %v 次", count)
}
