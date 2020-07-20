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

}
