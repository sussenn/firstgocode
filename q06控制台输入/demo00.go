package main

import (
	"fmt"
)

//控制台输入 scanln 和 scanf
func main() {
	var name string
	var age int
	var sex string

	//方式一 Scanln获取
	fmt.Println("输入姓名:")
	fmt.Scanln(&name)
	fmt.Println("输入年龄:")
	fmt.Scanln(&age)
	fmt.Println("输入性别:")
	fmt.Scanln(&sex)

	fmt.Printf("最终的结果打印: 姓名-%v, 年龄-%v, 性别-%v\n", name, age, sex)

	//方式二 Scanf
	var product string
	var num int
	var brand string
	fmt.Println("请输入商品, 数量, 品牌 [使用空格隔开]")
	fmt.Scanf("%s %d %s", &product, &num, &brand)
	fmt.Printf("最终的结果打印: 商品-%v, 数量-%v, 品牌-%v\n", product, num, brand)
}
