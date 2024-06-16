package main
import "fmt"
func main(){
	cityMap := make(map[string]string)
	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	//遍历
	for key,value := range cityMap{
		fmt.Println("key = ",key)
		fmt.Println("value = ",value)
	}
	//删除
	delete(cityMap,"China")
	//修改
	cityMap["USA"] = "DC"

}