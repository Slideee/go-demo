package main

import "fmt"

func main() {
	var command = "go inside"
	switch command {
	case "go east":
		fmt.Println("you head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("you find yourself in a dimly lit cavern.")
		fallthrough
	case "read sign":
		fmt.Println("the sign reads 'No Minors'")
	default:
		fmt.Println("quite")
	}
}
