package main

import (
	"fmt"
	"strconv"
)

//string类型转换其他类型
func main() {
	//bool
	var str string = "ture"
	var b bool

	//入参的位数必须和返回值位数一致,	入参int64,返回值也必须是int64
	//此函数返回2个值, _表示忽略此值
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b 的类型 = %T, b = %v\n", b, b)

	//int
	var str2 string = "123456"
	var i int64
	//p1:传参字符串  p2:10进制  p3:0表示int, 64表示int64
	i, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("i 的类型 = %T, i = %v\n", i, i)

	//double
	var str3 string = "12.3456"
	var flt float64
	flt, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("flt 的类型 = %T, flt = %v\n", flt, flt)

}
