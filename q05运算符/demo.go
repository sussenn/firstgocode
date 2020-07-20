package main

import (
	"fmt"
)

//运算符
func main() {
	var num int = 10
	//整数相除, 舍弃小数位 // num1 = 3
	num1 := num / 3
	fmt.Println("num1 = ", num1)
	//只有 i++, i--   没有 ++i, --i

	// && 短路与		// || 短路非
	// c += a	即 c = c+a
	// c <<= 2  即 c = C << 2 左移2位
	// 左移: 如 3左移3位 3 << 3 = 3 * 2^3 = 24

}
