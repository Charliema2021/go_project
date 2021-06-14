package main

import "fmt"

// 申明一个接口~
type Usb interface {
	Start() // 高内聚低耦合~
	Stop()
}
type Usb2 interface { //  golang中接口中的方法不用管这个名字，只要方法对就行~
	Start() // 高内聚低耦合~
	Stop()
}

type Phone struct {
}
type Camera struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作~~~~")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作~~~~")
}

func (ca Camera) Start() {
	fmt.Println("相机开始工作~~~~")
}
func (ca Camera) Stop() {
	fmt.Println("相机停止工作~~~~")
}

type Computer struct {
}

// 电脑有自己单独的方法，可以接受usb这个接口类型
func (c Computer) Working(u Usb) {

	// 通过usb接口变量来调用start和stop方法
	u.Start()
	u.Stop()

}

func main() {

	var ca = Camera{}
	var p = Phone{}
	var co = Computer{}

	// 使用~
	co.Working(ca) // 把p给你~~
	co.Working(p)

}
