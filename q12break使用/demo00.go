package main

import (
	"fmt"
	"math/rand"
	"time"
)

//for增强循环
func main() {
	var num int = 0
	for {
		//示例: 随机生成 1-100数字, 如果随机数是99则跳出循环
		//time.Now().UnixNano() 返回一个1970-00-00 00:00:00 至今的秒数
		rand.Seed(time.Now().UnixNano())
		//随机生成 1- 100 的整数
		n := rand.Intn(100) + 1
		fmt.Println("n = ", n)
		num++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共用了 ", num)

	//"变量名 + : "    即可作为一个标签
	//设置2个标签 myLable, 标签配合break 可以控制 break跳出范围(即决定在第一层for循环跳出,还是第二层for跳出)
	// myLable2:
	for i := 0; i < 4; i++ {
	myLable1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break myLable1 //外层循环了i=4次,因为跳出的是内层循环
				// break myLable2	//打印j=0 j=1 只循环了2次,便结束了
			}
			fmt.Println("j = ", j)
		}
	}

}
