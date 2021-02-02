package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	// 1、len
	// golang统一用utf-8，ASCII码都是1个字节，汉字是2个字节
	str := "hello"
	fmt.Println("str len=", len(str)) // 按照字节返回的~ 5
	str2 := "hello北"
	fmt.Println("str2 len=", len(str2)) // 按照字节返回的~ 8

	// 2、字符串遍历,str3[index]
	// 引入切片，因为含有中文字符~~~
	str3 := "hello北京" // 他的长度是11个字节
	//	r :=[]rune(str3) // 含有中文字符的给他转成切片~~！~！
	fmt.Println("str3 len=  ", len(str3)) // 11个字节~ 是utf-8的字节数
	//  两种获取字符串长度的方法~
	//	（1） golang中的unicode/utf8包提供了用utf-8获取长度的方法
	//	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
	fmt.Println("str3~ len=  ", utf8.RuneCountInString(str3))

	//   (2)  通过rune类型处理unicode字符
	//    fmt.Println("rune:", len([]rune(str)))

	fmt.Println("str3！ len=", len([]rune(str3)))

	for key, value := range str3 {
		fmt.Printf("key:%d value:%c \n", key, value)
	}

	/*
		for i:=0;i<=utf8.RuneCountInString(str3);i++{
			fmt.Printf("字符是=%c \n",r[i])
		} */

	/* 	for i:=0;i<len(str3);i++{
	fmt.Printf("字符是=%c  \n",r[i]) */
	// 报错 寻找原因：panic: runtime error: index out of range [7] with length 7
	//		实际遍历只有7个字符~

	// 3、字符串转整数 n,err ：= strconv.Atoi()

	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Printf("转换结果：%d；数据类型是：%T \n", n, n)
	}

	// 4、整数转字符串~ str = strconv.Itoa(12345)
	str4 := strconv.Itoa(12345)
	fmt.Printf("字符串是：%v, 格式是：%T \n", str4, str4)

	// *5、字符串转换成byte~~~

	var str5 = []byte("hello word")
	fmt.Printf("byte=%v \n", str5)

	// 6、数字转换成字符串~

	str6 := string([]byte{97, 98, 99})
	fmt.Println(str6)

	// 7、10进制转换成2、8、16进制   str= strconv.FormatInt(源数字，转换进制)

	str7 := strconv.FormatInt(100, 2)
	fmt.Printf("100对应的2进制数字：%v \n", str7)
	str7_1 := strconv.FormatInt(100, 16)
	fmt.Printf("100对应的16进制数字：%v \n", str7_1)

	// 8、判断字符串中是否含有子串~ string.Contains("原字符串"，"目标字符串"）bool

	b := strings.Contains("seafood", "food")
	fmt.Printf("b=%v \n", b)

	// 9、统计一个字符串中有几个相同的子串~   strings.Count("原字符串","关键词")

	num := strings.Count("dalskjflkadsjfieedzxjdfiesndkxdewewwss", "e")
	fmt.Printf("字符串中共有 %v个e \n", num)

	// 10、不区分大小写的字符串比较~  strings.EqualFold("ni","wo") bool

	c1 := strings.EqualFold("Abc", "ABC") //
	fmt.Printf("dubijieguowei:%v \n", c1)
	fmt.Println("结果", "Abc" == "ABC")

	// 11、返回子串在字符串的第几个位置~  ，如果没有返回-1！
	// 		strings.Index("dalkjeiakdjfiea;jdf","f")

	index := strings.Index("fdalkjeiakdjfiea;jdf", "f")      //！！！！！！ 他找到是第一次出现的位置！！！！！！ 而且计数是从0开始的 0 ，1,2,3,4,5...
	Lindex := strings.LastIndex("fdalkjeiakdjfiea;jdf", "f") // 最后一次出现的位置！
	fmt.Println("f出现在第", index)
	fmt.Printf("字符串长度为%v，f最后出现在%v \n", len("fdalkjeiakdjfieajdf"), Lindex)

	// 12、 替换字符~  strings.Replace("目标字符串"，"要替换的部分"，"替换成什么"，替换几个)
	//    	 当n=1时，顺序替换第一个，当n=-1时，全部替换
	str12 := "wo ai beijing 天安门"
	rstr12 := strings.Replace(str12, "wo", "ni", 1)
	fmt.Println("str12=", str12)
	fmt.Println("str12=", rstr12)

	// 13、将目标字符串儿以标识符分割，然后变成一个！！！字符串！！~  strings.Split(“目标字符串”，“标识符”)
	str13 := "nihao,woshi,nide,haopengyou"
	strArr := strings.Split(str13, ",") //以逗号为标识符，将目标字符串分割为新的数组
	fmt.Println(str13)
	//  strArr拿到的是一个字符串的值拷贝，实际字符串并没有被改变

	for key, value := range strArr { // 数组的遍历~ 美滋滋~
		fmt.Printf("key:%d,value=%v \n", key, value)
	}

	// 14、字母的大小写转换 strings.ToLower(目标字符串)
	// tolower小写  touper 大写

	str14 := strings.ToUpper(str13)
	fmt.Println("str13=", str13)
	fmt.Println("str14=", str14)

	// 15、处理一下字符串两边的空格~ strings.TrimSpace()

	str15 := strings.TrimSpace("  dafdsff   dfsf   ")
	fmt.Println("str15=", str15)

	// 16、去掉左右两边的字符~
	// 去掉左右的字符也有相应的命令！ trimleft、trimright~

	str16 := strings.Trim("!*hhaha,ni hao! ", "* !") // "字符直接罗列就行，不用符号隔开"  顺序也没有影响！
	fmt.Println("str16=", str16)
}
