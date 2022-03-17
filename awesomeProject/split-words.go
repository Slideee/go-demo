package main

import (
	"fmt"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go source2(c0)
	go filter2(c0, c1)
	printV2(c1)
}
func source2(c chan string) {
	for _, item := range []string{"apple is bad"} {
		c <- item
	}
	close(c)
}

func filter2(c chan string, down chan string) {
	for item := range c {
		for _, word := range strings.Fields(item) {
			down <- word
		}
	}
	close(down)
}

func printV2(c chan string) {
	for item := range c {
		fmt.Println(item)
	}
}
