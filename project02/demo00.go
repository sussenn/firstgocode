package main

import "fmt"

func main() {
	//go的转义字符
	// \t Tab键 实现对齐功能
	// \n 换行
	// \\ 一个斜杠
	// \" 一个双引号
	// \r 一个回车

	//cmd(DOS)窗口执行 gofmt -w xxx.go 可格式化.go文件
	//---------------------------------------------
	//DOS 常用命令
	//cd /d f:		进入f盘
	//md xxxFile	当前路径创建文件夹
	//rd xxxFile	删除文件夹
	//rd /q/s xxxFile		删除此目录所有文件
	//dir 			列出当前目录所有文件夹和文件
	//echo hello > d:\aaa\abc.txt 	创建abc.txt并写入hello
	//copy abc.txt f:\aaa			复制abc.txt到aaa文件夹下
	//copy abc.txt f:\aaa\ok.txt	复制abc.txt到aaa文件夹下,并重命名为ok.txt
	//move abc.txt e:\aaa			剪切abc.txt到aaa文件夹下
	//del abc.txt					删除abc.txt
	//cls 	清屏		exit	退出DOS

	fmt.Println("hello")
}
