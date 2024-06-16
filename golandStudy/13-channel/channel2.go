package main

import (
	"fmt"
	"time"
)

func main(){
	c:=make(chan int,3)//带缓存的channel
	fmt.Println("len(c) = ",len(c),"cap(c) = ",cap(c))
	go func() {
		defer fmt.Println("子进程结束")
		//当channel满了的时候会阻塞
		for i:=0;i<3;i++{
			c <- i
			fmt.Println("子进程正在运行，发送的元素 = ",i,"len(c) = ",len(c),"cap(c) = ",cap(c))
		}
	}()
	time.Sleep(2*time.Second)
	//带缓存的channel即使值没被读取，传值的进程也不会被阻塞；但当channel为空，读取channel里面的值会阻塞传值进程
	for i:=0;i<3;i++{
		num := <- c
		fmt.Println("父进程正在运行，接受的元素 = ",num)
	}
	fmt.Println("父进程结束")
}


