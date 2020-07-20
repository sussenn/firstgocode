package main

import "fmt"

func main() {
	//go build -o myhello.exe hello.go  //指定编译后.exe命名
	var i int = 10
	a := "我"
	fmt.Println("hello")
	fmt.Println("world")
	fmt.Printf("i = %d", i)
	fmt.Printf("\na=" + a)
}
