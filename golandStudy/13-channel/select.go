package main 
import "fmt"

func fibonacii(c,quit chan int){
	x,y := 1,1
	for{
		fmt.Println("select")
		//select每个case之后的条件只会阻塞对应的代码
		select {
		//不带缓存的channel发送的第一个数据没被接受再往其发送数据时，进程会被阻塞
		
		case c <- x:
			//如果c可写，执行下面的代码
			x = y
			y = y+x
			fmt.Println("case运行中")
		case <-quit:
			//如果quit可读，执行下面的代码
			fmt.Println("quit")
			return
		}
	}
	
}

func main(){
	c := make(chan int)
	quit := make(chan int)

	go func (){
		for i:=1;i<6;i++{
			fmt.Println(<-c)
		}
		fmt.Println("quit init")
		quit <- 0
	}()
	fibonacii(c,quit)
}