package main

import (
	"fmt"
)

//byte 和 string 类型 的打印
func main() {
	//byte
	var by1 byte = 'a'
	fmt.Println("by1实际值=", by1)
	fmt.Printf("by1输出值=%c", by1)

	//string ---------------------------------------------
	//1 string 类型一旦赋值就不能改变
	var str1 string = "hello"
	// str1[0] = "a"			//err
	fmt.Println("str1=", str1)
	//2 反引号里的内容不识别转义字符 如 \n
	var str2 string = `hello\n world`
	fmt.Println("str2=", str2)
}
