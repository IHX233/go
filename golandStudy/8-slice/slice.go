package main
import "fmt"
func printArray(myArray []int){
	for _,value:= range myArray{
		fmt.Println("value = ",value)
	}
	//所以这里可以修改原数组
	myArray[0] = 100
}
func main(){
	myArray := []int{1,2,3,4} //slice 
	//传入的是地址的指针 不传指针的情况下只有slice和map是传地址
	printArray(myArray)
	fmt.Println(myArray[0])
}