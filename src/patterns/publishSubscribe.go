package main

import "sync"

type PubSub struct {
	mu     sync.Mutex
	subs   map[string][]chan string
	closed bool
}

func NewPubSub() *PubSub {
	ps := &PubSub{}
	ps.subs = make(map[string][]chan string)
	return ps
}

func (ps *PubSub) Subscribe(topic string, ch chan string) {
	ps.mu.Lock()
	ps.subs[topic] = append(ps.subs[topic], ch)
	ps.mu.Unlock()
}

func (ps *PubSub) Publish(topic, msg string) {
	ps.mu.Lock()
	for _, ch := range ps.subs[topic] {
		ch <- msg
	}
	ps.mu.Unlock()
}
