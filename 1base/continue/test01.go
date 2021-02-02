package main

import "fmt"

//某人有100000元，没经过一次路口，需要缴费，当现金>50000时，交5%，当现金<=50000时，交钱1000，计算可以路过多少个路口

func main() {

	var corner int
	var money float64 = 100000

	for {
		if money > 50000 {
			money = money - money*0.05
			corner++
			continue
		}
		money = money - 1000
		corner++
		if money <= 0 {
			break
		}
	}
	fmt.Printf("10w元最懂能通过%v个路口", corner)

}
