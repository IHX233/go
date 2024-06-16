package main
import "fmt"
func main(){
	s:=[]int{1,2,3}
	s1:=s[0:2]//[1,2]
	//浅拷贝
	fmt.Println(s1)
	s1[0]=100
	fmt.Println(s)
	fmt.Println(s1)
	//深拷贝
	s2:=make([]int,3)
	copy(s2,s)
	s2[1]=100
	fmt.Println(s)
	fmt.Println(s2)
}