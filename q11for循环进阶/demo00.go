package main

import (
	"fmt"
)

//for增强循环
func main() {
	//1. 普通遍历
	var str string = "hello,world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	//2.增强for循环
	str2 := "sussenn-zhou"
	for index, st := range str2 {
		fmt.Printf("index = %d, st = %c \n", index, st)
	}

	//for 循环是按照字节来遍历的
	str3 := "hello world,杭州"
	str4 := []rune(str3) //把str转为切片
	for i := 0; i < len(str4); i++ {
		fmt.Printf("%c \n", str4[i])
	}

	//for-range遍历 是按照字符方式遍历
	str5 := "wode天啊"
	for _, v := range str5 {
		fmt.Printf("%c \n", v)
	}

}
