package main
import "fmt"
func main(){
	//第二个参数为长度，第三个参数为容量
	var numbers = make([]int,3,5)
	fmt.Printf("len = %d,cap = %d,slice=%v\n",len(numbers),cap(numbers),numbers)
	numbers = append(numbers,1)
	fmt.Printf("len = %d,cap = %d,slice=%v\n",len(numbers),cap(numbers),numbers)
	numbers = append(numbers,2)
	fmt.Printf("len = %d,cap = %d,slice=%v\n",len(numbers),cap(numbers),numbers)
	//当追加之后的切片长度大于当前容量，那么容量会在当前的容量加上初始容量
	numbers = append(numbers,3)
	fmt.Printf("len = %d,cap = %d,slice=%v\n",len(numbers),cap(numbers),numbers)
	//未定义容量的切片的容量会等于切片的长度
	numbers2 := make([]int,3)
	fmt.Printf("len = %d,cap = %d,slice=%v\n",len(numbers2),cap(numbers2),numbers2)
}