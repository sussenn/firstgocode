package main

import (
	"fmt"
)

//continue 跳出本次循环
func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			//跳出本次循环, 即不打印j=2
			if j == 2 {
				continue
			}
			fmt.Println("j = ", j)
		}
	}

	//示例: 只打印 20以内的奇数
	for n := 0; n < 20; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("奇数:", n)
	}

}
