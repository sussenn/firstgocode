package main

import (
	"fmt"
)

//for循环 语法
func main() {
	//1.写法1
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world i=", i)
	}

	//2.写法2
	j := 1
	for j <= 10 {
		fmt.Println("hello go j:", j)
		j++
	}

	//3.写法3	即无限for循环 for ; ; { //... }
	k := 1
	for {
		if k <= 5 {
			fmt.Println("hello for k-", k)
		} else {
			break
		}
		k++
	}

}
