package main
import (
	"fmt"
	"time"
	"errors"
	
)
//  golang中的错误处理机制
//  panic 来了~~~
func test(){

	//捕获异常
	defer func(){
		err := recover() // 内置函数可以直接捕获异常
		if err!=nil {    // 说明有异常
			fmt.Println("error=",err) // 打印错误
			// 同时在这里可以加入 提示~邮件发送等等功能
		}

	}()
	num1:= 10
	num2:=0
	res :=num1/num2    // 分母为0无法正常运算
	fmt.Println("result=",res)

}

//  函数去读取配置文件信息，
//  如果文件名不正确，就返回一个自定义的错误

func readFile(name string)(err error){
	if name=="config.exe"{
		//	输入正确
		return nil

	}else{
		return errors.New("get file failed!~~")
	}

}

func test02(){
	
	res :=readFile("config.exe")
	if res!=nil {
		// 如果出现错误，输出这个err并终止程序。
		panic(res) 
	}else{
		fmt.Println("程序继续执行...")
		time.Sleep(time.Second)
	}


}










func main(){


//  发生错误后捕获该错误，并进行处理，保证程序继续执行！
//  发现错误后，给管理员发送预警
//  defer panic recover
// ----------------------------------------------
// 使用refer+recover 来处理异常

	/* for {
		fmt.Println("继续执行中。。。")
		time.Sleep(time.Second)
	}
 */
	// 测试自定义错误
	test02()
}