package main

import "fmt"

func main() {
	const hours = 24
	var speed = 100800
	var distance = 96300000
	fmt.Println(distance/speed/hours, "days")
}
