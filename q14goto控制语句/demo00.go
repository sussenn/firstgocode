package main

import (
	"fmt"
)

//goto 跳转控制 	[不推荐使用! 逻辑容易混乱]
func main() {
	var n int = 20
	//如果n>10 则跳转到 label1标签 开始执行
	if n > 10 {
		goto label1
	}
	fmt.Println("go1")
	fmt.Println("go2")
label1:
	fmt.Println("go3")

}
