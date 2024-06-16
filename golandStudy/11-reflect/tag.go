package main
import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"gf"`
	Sex string `info:"sex"`
}

func FindTag(str interface{}){
	t := reflect.TypeOf(str).Elem()
	for i:=0; i<t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:",taginfo,"doc:",tagdoc)
	}
}

func main(){ 
	var re resume
	FindTag(&re)
}