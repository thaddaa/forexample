package main

import "fmt"

func main() {
	fmt.Println("tes")
	fof()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func fof() {
	fmt.Println("in foo")
}
