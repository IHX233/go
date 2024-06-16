package main
import "fmt"
func main(){
	//第一种声明方式
	//声明map1是一种map类型，key为stringvalue也为string
	var myMap1 map[string]string
	if myMap1 == nil{
		fmt.Println("myMap1 是个空map")
	}
	//在使用map前需要先为其分配空间
	myMap1 = make(map[string]string,10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"
	fmt.Println(myMap1)

	//第二种声明方式
	myMap2 := make(map[string]string)
	myMap2["one"] = "java"
	myMap2["two"] = "c++"
	myMap2["three"] = "python"
	fmt.Println(myMap2)


	//第三种声明方式
	myMap3 := map[string]string{
		"one":"php",
		"two":"js",
		"three":"go",
	}
	fmt.Println(myMap3)
}
