package main
import "fmt"

//返回一个值
func fool(a string,b string) int{
		fmt.Println("a = ",a)
		fmt.Println("b = ",b)
		c:= 100
		return c
}

//返回多个值
func fool2(a string,b string) (int,int){
		fmt.Println("a = ",a)
		fmt.Println("b = ",b)
		return 666,777
}
//返回多个有形参的值
func fool3(a string,b string) (r1 int,r2 int){
		fmt.Println("a = ",a)
		fmt.Println("b = ",b)
		r1 = 1000
		r2 = 2000
		return
}
//返回多个有形参切且同类型的值
func fool4(a string,b string) (r1 ,r2 int){
		fmt.Println("a = ",a)
		fmt.Println("b = ",b)
		r1 = 3000
		r2 = 4000
		return
}

func main(){
	c :=fool("a","b")
	fmt.Println("c = ",c)
	d,e :=fool2("a","b")
	fmt.Println("d = ",d)
	fmt.Println("e = ",e)
	f,g :=fool3("a","b")
	fmt.Println("f = ",f)
	fmt.Println("g = ",g)
	h,i :=fool4("a","b")
	fmt.Println("h = ",h)
	fmt.Println("i = ",i)
}