package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Philosopher struct {
	name           string
	leftChopstick  *sync.Mutex
	rightChopstick *sync.Mutex
}

func (p *Philosopher) run() {
	for {
		fmt.Printf("%d\t%s: Thinking\n", time.Now().UnixNano(), p.name)
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Printf("%d\t%s: Eating\n", time.Now().UnixNano(), p.name)
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
	}
}

const COUNT = 5

func main() {
	var chopsticks [COUNT]*sync.Mutex
	for i := 0; i < COUNT; i++ {
		chopsticks[i] = &sync.Mutex{}
	}

	for i := 0; i < COUNT-1; i++ {
		go (&Philosopher{fmt.Sprintf("Number %d", i), chopsticks[i+1], chopsticks[i]}).run()
	}
	go (&Philosopher{fmt.Sprintf("Number %d", COUNT-1), chopsticks[COUNT-1], chopsticks[0]}).run()

	time.Sleep(time.Second)
	os.Exit(0)
}
