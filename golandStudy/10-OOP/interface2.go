package main
import "fmt"
//interface{}是万能数据类型
func myFunc(arg interface{}){
	fmt.Printf("myFunc is called")
	fmt.Println(arg)
	//类型断言
	value,ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	}else{
		fmt.Println("arg is string,value = ",value)
		fmt.Printf("value type is %T\n",value)
	}

}
func main(){
	myFunc("sxj")
}