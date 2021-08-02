package main

import "fmt"

func main() {
	defer fmt.Println("This will be printed immediately when panic occurred!")
	fmt.Println("Get ready!")
	panic("LET'S PANIC!!!")
}
