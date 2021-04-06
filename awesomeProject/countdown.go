package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//countdown()
	//infinity()
	//rocket()
	guess()
}

func countdown() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
}

func infinity() {
	var degrees = 10
	for {
		fmt.Println(degrees)
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

func rocket() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		if rand.Intn(100) == 0 {
			break
		}
		count--
	}
	if count == 1 {
		fmt.Println("success!")
	} else {
		fmt.Println("failed!")
	}
}

func guess() {
	var answer = 55
	var guess = 0
	for answer != guess {
		guess = rand.Intn(100)
		fmt.Println(guess)
		if guess >= answer {
			fmt.Println("big than")
		} else {
			fmt.Println("less than")
		}
	}
}
