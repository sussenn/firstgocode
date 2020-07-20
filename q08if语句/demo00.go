package main

import (
	"fmt"
)

//if 语法
func main() {
	var age int = 5
	//
	if age > 10 {
		fmt.Printf("年龄大 age = %d\n", age)
	} else {
		fmt.Printf("年龄小 age = %d\n", age)
	}
}
