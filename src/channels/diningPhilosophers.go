package main

import (
	"fmt"
	"time"
)

func run(name string, leftChopstick, rightChopstick chan struct{}) {
	for {
		fmt.Printf("%d\t%s: Thinking\n", time.Now().UnixNano(), name)
		<-rightChopstick
		<-leftChopstick
		fmt.Printf("%d\t%s: Eating\n", time.Now().UnixNano(), name)
		leftChopstick <- struct{}{}
		rightChopstick <- struct{}{}
	}
}

const COUNT = 5

func main() {
	var chopsticks [COUNT]chan struct{}
	for i := 0; i < COUNT; i++ {
		chopsticks[i] = make(chan struct{}, 1)
		chopsticks[i] <- struct{}{}
	}

	for i := 0; i < COUNT-1; i++ {
		go run(fmt.Sprintf("Number %d", i), chopsticks[i+1], chopsticks[i])
	}
	go run(fmt.Sprintf("Number %d", COUNT-1), chopsticks[COUNT-1], chopsticks[0])

	time.Sleep(time.Second)
}
