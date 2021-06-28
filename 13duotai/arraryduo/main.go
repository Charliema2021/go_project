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
	Name string
}
type Camera struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "手机开始工作~~~~")
}
func (p Phone) Stop() {
	fmt.Println(p.Name, "手机停止工作~~~~")
}

func (ca Camera) Start() {
	fmt.Println(ca.Name, "相机开始工作~~~~")
}
func (ca Camera) Stop() {
	fmt.Println(ca.Name, "相机停止工作~~~~")
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

	// 多态数组~ 通过数组，将不同的结构体存在同一个数组内~~~实现多态的状况~~
	var usbArr [3]Usb
	var co = Computer{}

	usbArr[0] = Phone{"apple"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"nikon"}

	//fmt.Println(usbArr)
	co.Working(usbArr[0])
	co.Working(usbArr[1])
	co.Working(usbArr[2])

	/* 	var ca = Camera{}
	   	var p = Phone{}
	   	var co = Computer{}

	   	// 使用~
	   	co.Working(ca) // 把p给你~~
	   	co.Working(p)
	*/
}
