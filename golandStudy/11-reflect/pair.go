package main
import "fmt"
func main(){
	var a string
	a = "gf"
	//pair<statictype:string,value:"gf">
	var allType interface{}
	allType = a
	//pair<type:string,value:"gf">
	str,_ := allType.(string)
	fmt.Println(str)
}