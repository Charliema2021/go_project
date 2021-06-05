package main

import "fmt"

type MethodUitls struct {
	Chang float64
	Kuan  float64
}

func (m *MethodUitls) print() {
	for i := 1.0; i <= m.Chang; i++ {
		for j := 1.0; j <= m.Kuan; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func main() {
	var m MethodUitls
	m.Chang = 5
	m.Kuan = 3

	m.print()

}
