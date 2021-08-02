package main

import "fmt"

func main() {
	fmt.Println("Get ready!")
	panic("LET'S PANIC!!!")
	fmt.Println("This will not be executed")
}
