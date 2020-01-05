package main

import "fmt"

//摄氏温度
type Celsius float64

//华氏温度
type Fahrenheit float64

func main() {
	fmt.Println(CToF(32))
	var a = Celsius(20)
	fmt.Println(a)
}

func FToC(fahrenheit Fahrenheit) Celsius {
	//类型转换
	return Celsius((fahrenheit - 32) * 5 / 9)
}

func CToF(celsius Celsius) Fahrenheit {
	//类型转化
	return Fahrenheit(celsius*9/5 + 32)
}

//为Celsius类型指定方法集
func (d Celsius) String() string {
	return fmt.Sprintf("%g 摄氏度", d)
}
