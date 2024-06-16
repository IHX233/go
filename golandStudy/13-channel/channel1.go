package main

import "fmt"

func main(){
	//定义一个channel
	c := make(chan int)
	go func (){
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine正在运行")
		c <- 666 //把666赋值给c
	}()
	//如果c里面的值一直不被读取，那么子进程会一直被阻塞
	num := <- c//从c中接受数据，并赋值给num
	fmt.Println("num = ",num)
	fmt.Println("main goroutine 结束")

}