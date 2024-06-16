package main
import (
	"fmt"
)
var ga int = 22
var gb = 33
func main(){
	var a int
	fmt.Println("a = ",a)
	fmt.Printf("type of a = %T\n",a)
	var b int = 100
	fmt.Println("b = ",b)
	var c  = 100
	fmt.Println("c = ",c)
	//:=只能函数体内使用
	e:=100
	fmt.Println("e = ",e)
	fmt.Println("ga = ",ga)
	fmt.Println("gb = ",gb)
	var x,y int = 100,222
	fmt.Println("x = ",x,"y = ",y)
	var(
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ",vv,"jj = ",jj)
}