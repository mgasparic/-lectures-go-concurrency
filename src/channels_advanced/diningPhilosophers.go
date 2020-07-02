package main

import (
	"fmt"
	"time"
)

func run(name string, leftChopstick, rightChopstick chan int) {
	for {
		fmt.Printf("%d\t%s: Thinking\n", time.Now().UnixNano(), name)
		<-leftChopstick
		<-rightChopstick
		fmt.Printf("%d\t%s: Eating\n", time.Now().UnixNano(), name)
		rightChopstick <-  0
		leftChopstick <- 0
	}
}

const COUNT = 5

func main() {
	var chopsticks [COUNT]chan int
	for i := 0; i < COUNT; i++ {
		chopsticks[i] = make(chan int)
		go func(ch chan<- int) {
			ch <- 0
		}(chopsticks[i])
	}

	for i := 0; i < COUNT-1; i++ {
		go run(fmt.Sprintf("Number %d", i), chopsticks[i+1], chopsticks[i])
	}
	go run(fmt.Sprintf("Number %d", COUNT-1), chopsticks[COUNT-1], chopsticks[0])

	time.Sleep(time.Second)
}
