package main()
import "fmt"



func BubbleOrder(arr *[]int){

	for i:=0;i<len(*arr)-1;i++{	           //循环次数控制：数组元素个数为N，则循环次数为N-1次。
	for j:=0;j<len(*arr)-1-i;j++{       //大小比较次数控制：每次循环后比较次数会发生变化，为N-1-i,因为每轮循环比较后就会有一个数字确定下来 
		if (*arr)[j]<(*arr)[j+1]{   	//进行多次比较~交换大小数字位置
			minNum := 0		//初始化一个变量接收中间数字，进行交换用
			minNum = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = minNum	
	}
	}
	}
}



func main(){
	arr:=[5]int{24,69,80,57,13}

	BubbleOrder(&arr)
}