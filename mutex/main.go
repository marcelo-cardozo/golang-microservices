package main

import (
	"fmt"
	"sync"
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *AtomicInt) Increase() {
	// locks the scope so that only one thread can accessed at the same time
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *AtomicInt) Decrease() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value--
}

func (a *AtomicInt) Value() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

var (
	counter = &AtomicInt{}
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
	fmt.Printf("counter: %d", counter.Value())
}

func updateCounter(wg *sync.WaitGroup) {
	counter.Increase()
	wg.Done()
}
