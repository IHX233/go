package  main
import "fmt"
func main(){
	//声明slice1是一个切片，并且初始化，默认值是1，2，3。长度是3
	slice1 := []int{1,2,3}
	fmt.Println("slice1 = ",slice1)
	//声明slice2是一个切片，但是没有给slice2分配空间
	var slice2 []int
	fmt.Println("slice2 = ",slice2)
	//给slice2三个空间，默认值0
	slice2 = make([]int,3)
	fmt.Println("slice2 = ",slice2)
	//声明一个slice3，同时分配三个空间，默认值为0
	var slice3 = make([]int,3)
	fmt.Println("slice3 = ",slice3)
	//声明一个slice4，同时分配三个空间，默认值为0,通过:=推导出slice4是个切片
	slice4 := make([]int,3)
	fmt.Println("slice4 = ",slice4)
}