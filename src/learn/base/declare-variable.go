package main

import (
	"fmt"
	"os"
)

var i int

func main() {
	a, err := os.Open("C:\\Users\\gww\\Desktop\\work\\项目模块\\精准打击v1.5（测试中）\\数据文档")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result is ", a)

	/*
	无法编译
	 */
	//a, err := os.Open("C:\\Users\\gww\\Desktop\\work\\项目模块\\精准打击v1.5（测试中）\\数据文档")

	a, err = os.Open("C:\\Users\\gww\\Desktop\\work\\项目模块\\精准打击v1.5（测试中）\\数据文档")
	fmt.Print(a)
}
