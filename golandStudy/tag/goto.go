package main

import (
	"fmt"
)

func test() {
	i := 1
	if i == 2 {
		fmt.Println(i)
	} else {
		goto END
	}
END:
	fmt.Println("end...")
}
func main() {
	test()
}
