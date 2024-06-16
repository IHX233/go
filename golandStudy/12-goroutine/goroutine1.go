package main
import (
	"time"
	"fmt"
)
//子goroutine
func newTask(){
	i:=0
	for {
		i++
		fmt.Printf("new goroutine :i = %d\n",i)
		time.Sleep(1*time.Second)
	}
}
//父goroutine
func main(){
	//创建一个go程 去执行newTask()流程
	go newTask()
	fmt.Println("main goroutine exit")
	//注释的for循环代码可以让父goroutine一直执行，如果fu父goroutine不执行了，子也会不执行
	// i:=0
	// for{
	// 	i++
	// 	fmt.Printf("new goroutine :i = %d\n",i)
	// 	time.Sleep(1*time.Second)
	// }
}