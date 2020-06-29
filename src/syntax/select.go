package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
L:
	for {
		select {
		case <-quit:
			break L
		case c <- x:
			x, y = y, x+y
		}
	}
}

func main() {
	c := make(chan int, 3)
	quit := make(chan int)
	go fibonacci(c, quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
	for {
		select {
		case number := <-c:
			fmt.Println(number)
		default:
			return
		}
	}
}
