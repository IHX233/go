package main
import "fmt"
//声明一种新的数据类型myint，是int的一个别名
type myint int
//定义一个结构体
type Book struct{
	title string
	auth string
}
func changeBook(book Book){
	//传递book的一个副本
	book.auth="666"
}
func changeBook1(book *Book){
	//传递指针
	book.auth="777"
}
func main(){
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhangsan"
	fmt.Printf("%v\n",book1)
	changeBook(book1)
	fmt.Printf("%v\n",book1)
	changeBook1(&book1)
	fmt.Printf("%v\n",book1)

}
