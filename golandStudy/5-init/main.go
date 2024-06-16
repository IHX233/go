package main

import(
	"golandStudy/5-init/selfLibrary1"
	"golandStudy/5-init/selfLibrary2"
	//匿名导包，可以不使用包里的方法或者变量而不报错
	//_"golandStudy/5-init/selfLibrary2"
	//可以像lib2.Lib2Test()来使用其中的变量或方法
	//lib2"golandStudy/5-init/selfLibrary2"
	//可以像Lib2Test()直接使用lib2里的变量或方法
	//."golandStudy/5-init/selfLibrary2"
	"fmt"
)
//main是入口方法
func main(){
	fmt.Println("main")
	selfLibrary1.Lib1Test()
	selfLibrary2.Lib2Test()
}