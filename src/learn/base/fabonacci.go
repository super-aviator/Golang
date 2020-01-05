package main

import "fmt"

func main() {
	fabonacci(2)
	fabonacci(3)
	fabonacci(4)
	fabonacci(5)
	fabonacci(6)
}

func fabonacci(i int) {
	var x, y = 0, 1
	fmt.Print(x, y)
	fmt.Print(" ")
	for b := 0; b < i; b++ {
		x, y = y, x+y
		fmt.Print(y)
		fmt.Print(" ")

	}
	fmt.Println()
}
