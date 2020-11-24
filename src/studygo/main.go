package main

import "fmt"

//函数外面只能放一些声明

//声明变量 可以批量申明
var (
	name string
	age  int
	isOk bool
)

func main() {
	//go 中声明的变量必须被使用
	name = "david"
	age = 25
	isOk = true
	var temp string

	fmt.Printf("name:%s", name)
	fmt.Println(age) //ln会
	fmt.Print(isOk)

}
