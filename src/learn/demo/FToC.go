package main

import "fmt"

func main() {
	fmt.Printf("%g\n",fToC(212))
	fmt.Printf("%g\n",fToC(65))
}

func fToC(f float64) float64{
	return (f-32)*5/9
}
