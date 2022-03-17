package main

import "fmt"

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go source(c0)
	go filter(c0, c1)
	printV(c1)
}
func source(c chan string) {
	for _, item := range []string{"a", "b", "b", "c"} {
		c <- item
	}
	close(c)
}

func filter(c chan string, down chan string) {
	var prev string
	for item := range c {
		if item != prev {
			down <- item
			prev = item
		}
	}
	close(down)
}

func printV(c chan string) {
	for item := range c {
		fmt.Println(item)
	}
}
