package main  
  
import (  
	"fmt"  
	"time"  
)  
  
func main() {  
	ch1 := make(chan string)  
	ch2 := make(chan string)  
  
	go func() {  
		time.Sleep(1 * time.Second)  
		ch1 <- "one"  
	}()  
  
	go func() {  
		time.Sleep(1 * time.Second)  
		ch2 <- "two"  
	}()  
  
	select {  
	case msg1 := <-ch1:  
		fmt.Println("Received", msg1, "from ch1")  
	case msg2 := <-ch2:  
		fmt.Println("Received", msg2, "from ch2")  
	}  
  
	// 等待一段时间以确保goroutines有时间发送数据  
	time.Sleep(2 * time.Second)  
}