package main

import (
	"fmt"
)
   // 深入理解一下 引用类型这个意思
func main(){
	var arr [5]int =[5]int{10,20,30,40,50}
	slice:=arr[:]  // 等于取数组的整体哦
	// for循环遍历
	for i := 0; i < len(slice); i++{
		fmt.Printf("[%v]=[%v] \n",i,slice[i])
	}
	fmt.Println("------------------------")
	var slice02 []int = make([]int, 3) // 申明切片的第二个办法
	for i:=0;i<len(slice02);i++{
		slice02[i] = arr[i]
	}
	for i,v:=range slice02{
		fmt.Printf("[%v]=[%v] \n",i,v)
	} 
	fmt.Println("------------------------")
	//切片还是可以继续切片
	slice03 := slice[0:3]
	fmt.Println(slice03) //[10 20 30]
	fmt.Println("------------------------")
	slice03[1] =100       // 通过这种方式去改变
	fmt.Println(slice03)  //[10 100 30]
	fmt.Println(arr)      //[10 100 30 40 50]
// 因为slice是引用该空间，所以slice变化的时候，会影响arr发生相同变化。
}