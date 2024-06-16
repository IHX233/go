package main
import "fmt"
type Human struct {
	Name string
	Ad int
}
func (this *Human) Eat(){
	fmt.Println("human eat")
}
func (this *Human) Run(){
	fmt.Println("human run")
}
type SuperMan struct{
	Human //SuperMan继承Human类
	Level int
}
//重写方法
func (this *SuperMan) Eat(){
	fmt.Println("superman eat")
}
func (this *SuperMan) Fly(){
	fmt.Println("superman fly")
}
func main(){
	human := Human{Name:"gf",Ad:100}
	human.Eat()
	human.Run()
	superman := SuperMan{Human{Name:"sxj",Ad:100},88}
	superman.Run()
	superman.Eat()
	superman.Fly()
	fmt.Println(superman.Level)
}