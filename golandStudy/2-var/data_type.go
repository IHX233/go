package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 10
	fmt.Printf("%d \n", a) //十进制
	fmt.Printf("%b \n", a) //二进制
	fmt.Printf("%o \n", a) //八进制
	fmt.Printf("%x \n", a) //十六进制
	fmt.Printf("%c \n", a) //unicode

	fmt.Printf("%f \n", math.Pi)   //浮点数
	fmt.Printf("%.2f \n", math.Pi) //二位浮点数

	var b string = "ihx"
	fmt.Printf("%s \n", b) //字符串

	var c string = "free\tihx"
	fmt.Printf("%v \n", c) //转义字符串

	fmt.Printf("%T \n", math.Pi) //类型
}
