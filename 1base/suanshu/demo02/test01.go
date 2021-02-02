package main

import "fmt"

func main() {
	var day int = 97
	var workDay int
	var eday int
	workDay = day / 7
	eday = day % 7
	fmt.Printf("距离放假还有%d周，零%d天", workDay, eday)

}
