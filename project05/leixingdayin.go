package main

import (
	"fmt"
	"unsafe"
)

//数据类型 的打印
func main() {

	//int 类型				//具体类型,跟随系统. 此处实际是int64
	var n1 = 100
	fmt.Printf("n1数据类型是 %T", n1)

	var n2 int64 = 10
	fmt.Printf("\nn2 占用字节数是 %d", unsafe.Sizeof(n2))

	//浮点类型 单精度 float32	4字节 	/双精度 float64  8字节 类似double
	//单精度可能丢失尾部数据,造成精度损失
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1精度丢失=", num1, "num2=", num2)

	//浮点类型默认为 float64
	// var num3 = 1.1		//类型为float64
	// num4 := .123         //即0.123
	// num5 := 5.1234e2		//即5.1234 * 10^2 = 512.34
	// num5 := -5.1234e2	//即5.1234 / 10^2 = 0.051234
}
