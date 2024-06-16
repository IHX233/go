package main

import (
	"fmt"
)

func main() {
	//无缓存的channel接收空channel会阻塞，往没准备接收的channel中发送也会阻塞
	//var c chan int
	c := make(chan int)
	go func() {
		defer fmt.Println("子进程运行结束")
		for i := 0; i < 3; i++ {
			c <- i
		}
		//close可以关闭一个channel；关闭channel后不能再向channel发送数据，但是还可以继续读取数据
		close(c)
	}()
	for {
		//当channel无数据可读且已关闭，ok为false，否则为true
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}

	}
	//可以使用range来不断迭代操作channel
	// for data := range c {
	// 	fmt.Println(data, 222)
	// }
	fmt.Println("main finished")

}
