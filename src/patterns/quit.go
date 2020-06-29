package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, done <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case <-done:
				fmt.Println("go routine finished")
				return
			}
			time.Sleep(time.Duration(rand.Intn(1200)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	done := make(chan bool)
	c := boring("Joe", done)
	for i := rand.Intn(5); i >= 0; i-- {
		fmt.Println(<-c)
	}
	close(done)
	time.Sleep(2 * time.Second)
}
