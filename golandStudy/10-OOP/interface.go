package main
import "fmt"
//本质是一个指针
type AnimalIF interface{
	Sleep()
	GetColor() string
	GetType() string
}
//具体的类
type Cat struct{
	color string
}

func (this *Cat) Sleep(){
	fmt.Println("Cat is Sleep")
}
func (this *Cat) GetColor() string{
	return this.color
}
func (this *Cat) GetType() string{
	return "Cat"
}
func showAnimal(animal AnimalIF){
	animal.Sleep()
	fmt.Println("color = ",animal.GetColor())
	fmt.Println("type = ",animal.GetType())
}

//具体的类
type Dog struct{
	color string
	
}
func (this *Dog) Sleep(){
	fmt.Println("Dog is Sleep")
}
func (this *Dog) GetColor() string{
	return this.color
}
func (this *Dog) GetType() string{
	return "Dog"
}
func main(){
	// var animal AnimalIF
	// animal = &Cat{"Green"}
	// animal.Sleep()
	// animal.GetColor()
	// animal.GetType()
	// animal = &Dog("yellow")
	// animal.Sleep()
	animal := Cat{"Green"}
	showAnimal(&animal)
}

