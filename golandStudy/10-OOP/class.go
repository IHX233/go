package main
import "fmt"
type Hero struct {
	//类的属性首字母大写表示对外可见
	Name string
	Ad int
	Level int
}
// func (this Hero) Show(){
// 	fmt.Println("Name = ",this.Name)
// 	fmt.Println("Ad = ",this.Ad)
// 	fmt.Println("Level = ",this.Level)
// }
// func (this Hero) SetName(newName string){
// 	//this 是调用该方法的对象的一个副本
// 	this.Name = newName
// }
func (this *Hero) Show(){
	fmt.Println("Name = ",this.Name)
	fmt.Println("Ad = ",this.Ad)
	fmt.Println("Level = ",this.Level)
}
func (this *Hero) SetName(newName string){
	this.Name = newName
}
func main(){
	hero := Hero{Name:"sxj",Ad:100,Level:1}
	hero.Show()
	hero.SetName("gf")
	hero.Show()
}