package main
import "fmt"
func main(){
	var myArray1 [10] int
	myArray2 :=[10] int {1,2,3,4}
	for i:=0;i<len(myArray1);i++{
		fmt.Println(myArray1[i])
	}
	for index,value:=range myArray2{
		fmt.Println("index = ",index,"value = ",value)
	}
}