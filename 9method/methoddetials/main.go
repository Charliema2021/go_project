package main

import "fmt"

type integer int

func (i integer) print() {
	fmt.Println(i)
}

func (i *integer) sum() {

	*i = *i + 1

}

func main() {
	var i integer = 10
	(&i).sum()

	fmt.Println(i)
}
