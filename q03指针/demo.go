package main

import (
	"fmt"
)

//指针
func main() {
	var i int = 10
	// i 的地址值
	fmt.Println("i 的地址值 = ", &i)

	//指针类型
	//[注意]: 1.[指针类型 必须指向地址值 即&i]
	//		  2.指针类型 必须和指向地址值的数值类型一致. 即 ptr *int64 对应 i int64才行

	// ptr是指针类型, 指向i的地址值(内存空间)
	var ptr *int = &i
	// ptr = i的地址值 = 0xc0000120d0
	fmt.Printf("ptr = %v\n", ptr)

	// ptr实际也有一个自己真正的地址值(内存空间) 	[注意]: 该内存空间只能有一个值, 后来再赋值会覆盖前面的赋值
	// &prt 的地址值 =  0xc000006030
	fmt.Println("prt 的地址值 = ", &ptr)

	// 取出ptr指向空间对应的值 即取出i的数值
	// *ptr = 10
	fmt.Println("prt 的值 = ", *ptr)

}
