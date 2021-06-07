package personinfor

import "fmt"

type Personinfor struct {
	Name   string  // name字段可以随便访问~~就大写
	age    int     // age 不可以随意访问，所以小写~~~
	salary float64 // salary 保密也小写~~
}

func NewPerson(name string) *Personinfor {
	// 新建一个人名字~~~只有名字~~其他内容都系统默认~0
	return &Personinfor{
		Name: name,
	}

}

//为了访问age和salary所以需要对这两个函数进行专用的方法

func (p *Personinfor) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("输入的年龄有误~~")

	}

}

func (p *Personinfor) GetAge() int {
	return p.age
}

func (p *Personinfor) SetSalary(sal float64) {
	if sal > 3000 && sal <= 30000 {
		p.salary = sal
	} else {
		fmt.Println("输入的金额范围有误~~")

	}

}

func (p *Personinfor) GetSalary() float64 {
	return p.salary
}
