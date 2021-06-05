package main

import "fmt"

type cal struct {
	num1     float64
	num2     float64
	operator string
}

func (c cal) jisuan() float64 {
	res := 0.0
	switch c.operator {
	case "+":
		res = c.num1 + c.num2
	case "-":
		res = c.num1 * -c.num2
	case "*":
		res = c.num1 * c.num2
	case "/":
		res = c.num1 / c.num2
	default:
		fmt.Println("错误！")
	}
	return res
}
func main() {
	var c1 cal
	c1.num1 = 5
	c1.num2 = 10
	c1.operator = "^"

	fmt.Printf("%v %v %v =%2.2f  \n", c1.num1, c1.operator, c1.num2, c1.jisuan())

}
