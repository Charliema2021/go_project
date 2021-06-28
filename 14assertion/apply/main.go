// 改造一下这个程序～用类型确认给phone增加一个专属的功能叫calling
package main

import "fmt"

// 申明一个接口~
type Usb interface {
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

// 给手机一个特有的方式～
func (p Phone) Calling() {
	fmt.Println(p.Name, "手机呼叫中~~~~")
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
	//通过类型断言来判断传进来的u是不是phone，如果是就加新方法，如果不是继续运行～
	if u, ok := u.(Phone); ok {
		u.Calling()
	}
	u.Stop()

}

func main() {

	// 多态数组~ 通过数组，将不同的结构体存在同一个数组内~~~实现多态的状况~~
	var usbArr [3]Usb
	var co = Computer{}

	usbArr[0] = Phone{"apple"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"nikon"}
	//对数组进行遍历～
	for _, v := range usbArr {
		co.Working(v)
	}
}
