package main

import "fmt"

//此处定义全局变量	不能使用类型推导 :=
var a, b, c = 10, "字", 80.9
var (
	d = "hello"
	// e = 100
)

func main() {
	//1声明多个不同类型的变量
	// var a, b, c = 10,"字",80.9
	//2
	// a, b, c := 10,"字",80.9
	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d)

}
