package main

import (
	"fmt"
)
func swap(pa int,pb int){
	var temp int
	temp = pa
	pa = pb
	pb = temp
}
//此时pa、pb为传入参数的内存地址
func swap1(pa *int,pb *int){
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}
func main(){
	var a int = 10
	var b int = 20
	fmt.Println("a = ",a,"b = ",b)
	swap(a,b)
	fmt.Println("a = ",a,"b = ",b)
	//此时传入的为a，b的内存地址
	swap1(&a,&b)
	fmt.Println("a = ",a,"b = ",b)
}