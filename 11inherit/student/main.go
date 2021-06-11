package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	score    float64 // 学生成绩属于秘密不能随便的修改~~~
	identity string
}

type Pupil struct {
	Student
}

func (stu *Student) testing() {
	fmt.Printf("%v考试中~~~~~", stu.identity)
}

// 获取成绩的方法~~~~
func (stu *Student) GetPupilScoer() float64 {
	return stu.score
}

// 获取identify的方法
func (stu *Student) GetIdentify() string {
	return stu.identity
}

// 修改成绩的方法
func (stu *Student) SetScoer(score float64) {
	// 后期系统升级~~增加校验哦~~
	//如果是老师就可以修改成绩，如果不是老师就不能

	if stu.GetIdentify() != "teacher" {
		fmt.Println("你无权修改成绩！~~")
		return
	} else {
		stu.score = score
		fmt.Println("成绩修改成功~~~")
	}

}

// 显示一个学生的成绩~~查询功能
func (stu *Student) showInfor() {
	fmt.Printf("name=[%v],age=[%v],id=[%v],score=[%v]\n", stu.Name, stu.Age, stu.identity, stu.score)
}

func main() {
	var pupil01 = &Pupil{}
	pupil01.Student.Name = "matianqi"
	pupil01.Student.Age = 20
	pupil01.Student.score = 98.00
	pupil01.Student.identity = "pupil"

	//pupil01.Student.showInfor()
	pupil01.Student.SetScoer(89.00)
}
