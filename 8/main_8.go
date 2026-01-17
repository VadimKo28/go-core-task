package main

import (
	"sync"
	"fmt"
	"time"
)

type WaitGroup struct {
	mu      sync.Mutex
	counter int
	sem     chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sem: make(chan struct{}),
	}
}

func (wg *WaitGroup) Add(count int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.counter += count

	if wg.counter < 0 {
		panic("negative WaitGroup counter")
	}

	if wg.counter == 0 {
		close(wg.sem)
		wg.sem = make(chan struct{})
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	wg.mu.Lock()
	sem := wg.sem
	if wg.counter == 0 {
		wg.mu.Unlock()
		return
	}
	wg.mu.Unlock()

	<-sem
}


func main() {
	wg := NewWaitGroup()
	
	fmt.Println("Запуск 5 горутин...")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Горутина %d начала работу\n", id)
			time.Sleep(time.Duration(id+1) * time.Second)
			fmt.Printf("Горутина %d завершила работу\n", id)
		}(i)
	}
	
	fmt.Println("Ожидание завершения всех горутин...")
	wg.Wait()
	fmt.Println("Все горутины завершены!")
}
