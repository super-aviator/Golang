package main

import "fmt"

func main() {
	fmt.Println(gcd(9, 18))
	fmt.Println(gcd(42, 36))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
