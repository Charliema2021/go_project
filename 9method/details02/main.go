package main

import "fmt"

type stuInfor struct {
	Name string
	Age  int64
}

func (stu *stuInfor) string() string {
	str := fmt.Sprintf("姓名=[%v]，年龄=[%v]\n", stu.Name, stu.Age)
	return str

}

func main() {

	var stu1 stuInfor
	stu1.Name = "tom"
	stu1.Age = 22
	// 为啥调用不到这个方法呢~~~
	str := (&stu1).string()
	fmt.Sprintf(str)

}
