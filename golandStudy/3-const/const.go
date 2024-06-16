package main
import(
	"fmt"
)
//const 只能定义只读变量
const (
	//iota 默认从0开始，只能配合const使用
	a,b = iota+1,iota+2 //iota=0 a=iota+1 b=iota+2
	c,d                 //iota=1 c=iota+1 d=iota+2
	e,f                 //iota=2 e=iota+1 f=iota+2

	g,h = iota*2,iota*3 //iota=3 g=iota*2 h=iota*3
	i,k                 //iota=4 i=iota*2 k=iota*3
)
func main(){
	fmt.Println("a = ",a,"b = ",b)
	fmt.Println("c = ",c,"d = ",d)
	fmt.Println("e = ",e,"f = ",f)
	fmt.Println("g = ",g,"h = ",h)
	fmt.Println("i = ",i,"k = ",k)
}
