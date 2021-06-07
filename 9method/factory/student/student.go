package student




type student struct{
	Name  string
	Age  int
	score float64
	
}

//  写一个方法，传入数据，然后返回一个student
func NewStu(n string,a int,s float64) *student{

	return &student{
		Name : n,
		Age : a,
		score : s,   // 这里score首字母小写，在其他包就没法正常用，处理方式是给他单独一个方法
	}

}

func (stu *student)GetScore() float64{

	return stu.score 

}
