package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, res chan<- string) {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	res <- fmt.Sprintf("%s", msg)
}

func main() {
	res := make(chan string)
	go boring("Joe", res)
	go boring("Peter", res)
	go boring("Mia", res)
	for i := 0; i < 3; i++ {
		fmt.Printf("You say: %q\n", <-res)
	}
	fmt.Println("The chat is over. Too boring.")
}
