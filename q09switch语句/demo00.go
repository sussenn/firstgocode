package main

import (
	"fmt"
)

//switch 语法
func main() {
	var i int
	fmt.Println("请输入0-9数字")
	fmt.Scanf("%d", &i)

	// 1. case/switch后是一个表达式 即变量,常量,函数都可以
	// 2. switch后可不写条件,直接在case后进行条件判断
	switch i {
	case 0:
		fmt.Println("你是零")
	case 1:
		fmt.Println("你是一")
	default:
		fmt.Println("你是基佬")
	}

	// 穿透fallthrough	case语句块后增加fallthrough,则会继续执行下一个case
	var num int
	switch num {
	case 0:
		fmt.Println("零")
		fallthrough //会继续执行下一个case,默认只能穿透一层
	case 1:
		fmt.Println("一")
	default:
		fmt.Println("基佬")
	}
}
