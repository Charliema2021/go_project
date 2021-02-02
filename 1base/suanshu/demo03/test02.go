package main

import "fmt"

func main() {
	var htemperature int = 60
	var stemperature int

	stemperature = 5 / 9 * (htemperature - 100)
	fmt.Println("华氏温度：", htemperature, "对应的摄氏温度是：", stemperature)

}
