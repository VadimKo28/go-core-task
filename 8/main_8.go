package main

type Semaphore struct {
	C chan struct{}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.C
}

type SemaphoreWaitGroup struct {
	counter int
	sem     *Semaphore
	done    chan struct{}
}

func NewSemaphoreWaitGroup() *SemaphoreWaitGroup {
	wg := &SemaphoreWaitGroup{
		sem: &Semaphore{C: make(chan struct{}, 1),},
		done: make(chan struct{}),
	}
	wg.sem.Acquire()
	return wg
}

func (wg *SemaphoreWaitGroup) Add(delta int) {
	wg.sem.Acquire()
	
	wg.counter += delta
	
	if wg.counter < 0 {
		wg.sem.Release()
		panic("negative counter")
	}
	
	if wg.counter == 0 {
		select {
		case <-wg.done:
		default:
			close(wg.done)
		}
	} else {
		select {
		case <-wg.done:
			wg.done = make(chan struct{})
		default:
		}
	}
	
	wg.sem.Release()
}

func (wg *SemaphoreWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *SemaphoreWaitGroup) Wait() {
	<-wg.done
}

func main() {
	wg := NewSemaphoreWaitGroup()
	
	wg.Add(3)
	
	go func() {
		defer wg.Done()
		println("Goroutine 1 started")
		println("Goroutine 1 finished")
	}()
	
	go func() {
		defer wg.Done()
		println("Goroutine 2 started")
		println("Goroutine 2 finished")
	}()
	
	go func() {
		defer wg.Done()
		println("Goroutine 3 started")
		println("Goroutine 3 finished")
	}()
	
	println("Waiting for all goroutines to complete...")
	wg.Wait()
	println("All goroutines completed!")
	
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		println("Goroutine 4 started")
		println("Goroutine 4 finished")
	}()
	
	go func() {
		defer wg.Done()
		println("Goroutine 5 started")
		println("Goroutine 5 finished")
	}()
	
	println("Waiting for goroutines 4 and 5...")
	wg.Wait()
	println("All done!")
}
