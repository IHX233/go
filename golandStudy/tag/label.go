package main

import "fmt"

func test() {
MYLABEL:
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i > 5 {
			break MYLABEL
		}
	}

	fmt.Println("end...")
}
func main() {
	test()
}
