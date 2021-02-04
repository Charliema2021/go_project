// 二分法 显得让数列顺序排列，从小到大~！~！
package main

import "fmt"

// func change  先给数列变成从小到达排列

func erfenSerch(arr *[6]int,value,LeftIndex,RightIndex int){
	//退出条件
	if LeftIndex>RightIndex{
		fmt.Println("没有这个数字。")
		return
	}
	erfen :=(RightIndex+LeftIndex)/2    // int类型是向下取整，所以没问题。
	
	if (*arr)[erfen]>value{   //如果中间值小于value，则RightIndex=erfen
		//查找范围在左边，右边距左移减一
		RightIndex = erfen-1
		erfenSerch(arr,value,LeftIndex,RightIndex)
	}else if (*arr)[erfen]<value{
		//查找范围在右边，左边距右移加一
		LeftIndex =erfen+1
		erfenSerch(arr,value,LeftIndex,RightIndex)
	}else {
		fmt.Printf("找到了,是第%v个数:",erfen+1)
	}

}


func main(){
	arr:=[6]int{1,30,56,78,99,1314}

	// 递归调用，调自己
	erfenSerch(&arr,1314,0,len(arr)-1)

}