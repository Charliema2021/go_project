package main

import "fmt"


func visit(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
} 

func main() {
    s := []int{1, 2, 3}                          
    ss := s[1:]         // 从s中取出切片                                     
    ss = append(ss, 4)  // 给ss追加了空间,赋值4

    for func(_, v := range ss)int {
        v += 10
        // 遍历ss并对每个元素进行累加求和
        return v
    }
    //fmt.Println(v)
    for i := range ss {   // 遍历ss并给每个元素加个10
        ss[i] += 10
    }
    fmt.Println(s) //   123 
    fmt.Println(ss) //   123 
    visit(ss,func(value int)){
        fmt.Println(value)
    }
    

}