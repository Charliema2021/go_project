//1\简化写法的情况~
// 在go中继承1是会一层一层的验证的
package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}
type Bands struct {
	Name    string
	Address string
}

// 常用方法~~多重继承~！~~！~！~
type TV struct {
	Goods
	Bands
}

// 通过指针去调用，效率更高~~
type TV2 struct {
	*Goods
	*Bands
}

func main() {
	//快速给定结构体内变量的值~~比较好的地方是，变换变量顺序也没有关系~~
	tv01 := TV{
		Goods{
			Name:  "电视机",
			Price: 1000.99,
		},
		Bands{
			Name:    "xiaomi",
			Address: "北京",
		},
	}
	tv02 := TV2{
		&Goods{
			Name:  "电视机",
			Price: 1000.99,
		},
		&Bands{
			Name:    "三星",
			Address: "韩国",
		},
	}

	fmt.Printf("你选择的货物为%v，价格为%v,品牌是%v，产地是：%v \n", tv01.Goods.Name, tv01.Price, tv01.Bands.Name, tv01.Address)
	fmt.Println(tv02)                     // 直接返回来了 两个地址~~
	fmt.Println(*tv02.Goods, *tv02.Bands) // 取值~~

	//fmt.Printf("你选择的货物为%v，价格为%v,品牌是%v，产地是：%v \n", tv02.Goods.Name, tv02.Price, tv02.Bands.Name, tv02.Address)

}
