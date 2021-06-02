package cal

import "fmt"

// 自定义函数，主要格式func cal(形式参数 参数类型，形式参数 参数类型，形式参数 参数类型)（返回值列表，返回值列表）
func Cal(n1 float32, n2 float32, operater string) float32 {
	var result float32
	switch operater {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		result = n1 / n2
	default:
		fmt.Println("输入符号错误！")
		break
	}
	return result // 需要返回的结果
}

/* func main() {
	var n1 float32 = 1.2
	var n2 float32 = 1.8
	var operater string = "+"

	fmt.Println("请输入数字n1: ")
	fmt.Scanln(&n1)

	fmt.Println("请输入数字n2: ")
	fmt.Scanln(&n2)

	fmt.Println("请输入运算符: ")
	fmt.Scanln(&operater)
	mainResult := cal(n1, n2, operater) // cal（形式参数，形式参数）把形式参数传给cal方法，然后把返回结果result赋给mainResult
	fmt.Println("result=", mainResult)
}
*/
