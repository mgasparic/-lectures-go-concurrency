package main

import (
	"fmt"
	"sync"
	"time"
)

func inc() {
	mux.Lock()
	counter++
	mux.Unlock()
}

func dec() {
	mux.Lock()
	counter--
	mux.Unlock()
}

var (
	counter int
	mux     sync.Mutex
)

func main() {
	for i := 0; i < 500; i++ {
		go inc()
	}
	for i := 0; i < 200; i++ {
		go dec()
	}
	time.Sleep(time.Second)
	fmt.Println(counter)
}
