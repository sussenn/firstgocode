package main

import (
	"fmt"
	"strconv"
)

//类型转换
func main() {
	//int 类型转换
	var i int32 = 10
	//类似Java强转
	//注意: i 本质还是int32类型
	//大范围的类型向小范围类型强转时, 编译不会报错,但如果数值过大,超出向下转型类型的范围时, 结果失真.
	var n1 float32 = float32(i)
	fmt.Printf("n1=%v\n", n1)

	//=====================================
	//string 类型转换
	//1 使用fmt.Sprintf
	var str1 string
	str1 = fmt.Sprintf("%d", i)                  //%d double类型占位,即int类型
	fmt.Printf("str1类型=%T, str= %v", str1, str1) //%v varchar类型占位,即字符串类型
	//%t	布尔类型占位

	//2 使用strconv包
	var num int = 10
	var flt float64 = 10.12345
	var str2 string
	//int转string
	str2 = strconv.FormatInt(int64(num), 10)      //p1:必须是int64类型数值 p2:表示十进制
	fmt.Printf("str2类型=%T, str2= %q", str2, str2) //字符串类型推荐 %q 会用""进行包裹
	//int转string 2
	//传参必须是int类型
	str2 = strconv.Itoa(num)
	fmt.Printf("str2类型=%T, str2= %q", str2, str2)

	//float64转string
	//p1:传参数值 p2:输出格式,常用'f' p3:小数点后保留多少位 p4:表示传参float64
	str2 = strconv.FormatFloat(flt, 'f', 10, 64)
	fmt.Printf("str2类型=%T, str2= %q", str2, str2)

	//bool转string
	//str2 = strconv.FormatBool(布尔类型值)

}
