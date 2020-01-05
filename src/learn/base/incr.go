package main

import "fmt"

func main() {
	a := 1
	fmt.Println(incr(&a))
	fmt.Println(incr(&a))
}

func incr(a *int) int {
	*a++
	return *a
}
