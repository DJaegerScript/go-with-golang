package main

import "fmt"

func main() {
	defer fmt.Println("This will be printed last")
	defer fmt.Println("This will be printed second")
	defer fmt.Println("This will be printed first")
}
