package main

import "fmt"

/* 从简单到复杂
1、先打矩阵

2、调整for循环，让每层打出相应的*


3、打印空格的规律
  *			* 层数，空格，总层数-当前层数
 ***
*****

4、空心，条件格式，什么时候打印*，什么时候打印空格
  *
 * *
*****   层数=最大层数时候，全部打*
先打空格，打完之后打*（第一个打，最后一个打，index是1和2i-1）然后


*/

func main() {

	var level int // 总层数
	fmt.Println("请输入层数")
	fmt.Scanln(&level)
	// i 是层数
	for i := 1; i <= level; i++ {

		// k 是每层空格的个数
		for k := 1; k <= level-i; k++ { //空格数十总层数见当前层数
			fmt.Print(" ")
		}
		// j 是每层*个数
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == level {
				fmt.Print("*") // println是格式化输出，会换行，print是直接打印输出，不会换行
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println() // 层间换行
	}

}

//   *
//  * *
// *  **
//*  *  *
// 出现这个问题的原因分析~！~
