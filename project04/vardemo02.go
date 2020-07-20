package main

import "fmt"

//go 基本数据类型 说明
func main() {
	//同一变量不能重复定义赋值
	var i int = 10
	// var i int = 20	err
	// i := 30 			err

	// int8 的范围 即 -2^7 ~ 2^7 -1		//有8位,第一位是标识正负,余下7位存储数值. 最后减一是因为有+0000000,-0000000存在,只取-0,所以正数范围-1
	// int16 -2^15 ~ 2^15 -1

	// uint8 uint16 uint32 uint64 		//舍弃负值,只取正值范围

	//rune 								//等价于int32

	//var v byte = 'a'					//等价于uint8
	fmt.Println("i=", i)
}
