package main

//继承无法完成的工作～ 可以通过接口灵活完成～ 主要原因是在于
import "fmt"

// 先定义一个老猴子～有自己固有的方法～
type Monkey struct {
	Name string
}
type LittleMonkey struct {
	Monkey //继承monkey的基本方法～
}

func (mon *Monkey) climbing() {
	fmt.Println(mon.Name, "clinbing~~~~")
}

// 然后通过接口对小猴子的能力进行扩展～
// fly
type BirdAble interface {
	flying()
}

// swimming
type FishAble interface {
	swimming()
}

// littlemonkey 开始通过接口完成能力扩展～
// 学习飞行～～
func (Lmon *LittleMonkey) flying() {
	fmt.Println(Lmon.Name, "学会飞了～～")
}

// 学习游泳～～
func (Lmon *LittleMonkey) swimming() {
	fmt.Println(Lmon.Name, "学会飞了～～")
}

func main() {
	var lmon LittleMonkey
	lmon.Name = "Sunwukong"

	lmon.climbing()
	lmon.flying()
	lmon.swimming()
}
