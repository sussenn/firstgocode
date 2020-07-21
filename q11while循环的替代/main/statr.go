package main

import "fmt"

//没有 while和do..while 使用for替代==========================================
func main() {
	var n int = 5
	for {
		if n > 10 {
			break
		}
		fmt.Println("使用for if 替代while", n)
		n++
	}

	//do while 的替代
	var m int = 1
	for {
		fmt.Println("进来先执行此代码,再作判断", m)
		m++
		if m > 3 {
			break
		}
	}

	//多重循环=============================================================
	var totalLovel int = 20
	//i 表示层级
	for i := 0; i < totalLovel; i++ {
		//先打印空格
		for k := 1; k < totalLovel-i; k++ {
			fmt.Print(" ")
		}
		//j 表示每层多少个*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLovel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//打印九九乘法
	var num int = 9
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}
