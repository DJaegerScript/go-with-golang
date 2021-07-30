package main

import "fmt"

func main() {
	someone := 10
	someoneElse := &someone

	fmt.Println("Here, someone's value is", someone)
	fmt.Println("And someone else's value is", *someoneElse)
}
