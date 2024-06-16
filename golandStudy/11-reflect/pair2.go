package main

import (
	"fmt"
	"reflect"
)

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriterBook()
}
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("reading")
}
func (this *Book) WriterBook() {
	fmt.Println("writering")
}
func main() {
	//b:pair<type:Book,value:book{}地址>
	b := &Book{}
	//r:pair<type:,value:>
	var r Reader
	fmt.Println("type:", reflect.TypeOf(r))
	fmt.Println("value:", reflect.ValueOf(r))
	//r:pair<type:Book,value:book{}地址>
	r = b
	fmt.Println("type:", reflect.TypeOf(r))
	fmt.Println("value:", reflect.ValueOf(r))
	r.ReadBook()
	var w Writer
	fmt.Println("type:", reflect.TypeOf(w))
	fmt.Println("value:", reflect.ValueOf(w))
	w = r.(Writer) //此处为什么会断言成功，因为w，r的类型一致
	fmt.Println("type:", reflect.TypeOf(w))
	fmt.Println("value:", reflect.ValueOf(w))
	w.WriterBook()

}
