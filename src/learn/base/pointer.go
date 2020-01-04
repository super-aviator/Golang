package main

import "fmt"

func main() {
	var i = 12
	var p = &i
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

	//无法编译，只能对指针类型的值进行取值操作
	//fmt.Println(*i)
}
