package main

import "fmt"

func test(n1 int) { // 形参n1是从main里过来的，现在是10
	n1 = n1 + 1
	fmt.Println("test(n1)=", n1)
	// 函数执行完毕，操作系统会回收、销毁这个栈空间，n1=11 就没了
}

func getSum(n2 int32, n3 int32) int32 {

	sum := n2 + n3
	return sum

}

func main() {
	// 代码执行顺序是一句一句执行的，每个func都会在栈堆里有一个区域，用完回收
	var n1 int = 10 //1 变量声明

	test(n1) //2 把main里的n1传给test，然后执行test

	fmt.Println("main(n1)=", n1) // 3 打印n1,此时test里的n1回收，现在只有main里的n1=10

	sum := getSum(10, 20)

	fmt.Println("main(sum)=", sum)
}
