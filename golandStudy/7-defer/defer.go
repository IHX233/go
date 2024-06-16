package main
import "fmt"
//defer先入后出
func deferOrder(){
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
}
//return比defer先执行
func deferFunc() string{
	defer fmt.Println("defer")
	return "deferFunc"
}
func returnFunc() string{
	fmt.Println("return")
	return "returnFunc"
}
func deferAndReturn() string{
	defer deferFunc()
	return returnFunc()
}
func main(){
	deferOrder()
	deferAndReturn()
}