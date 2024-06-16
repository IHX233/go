package selfLibrary1

import(
	"fmt"
)
//大写的才是可见的
func Lib1Test(){
	fmt.Println("Lib1Test")
}
//init方法是包里的自执行方法，无需调用
func init(){
	fmt.Println("lib1")
}