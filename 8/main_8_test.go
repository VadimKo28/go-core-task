package main

import (
	"testing"
	"time"
)

func TestSemaphoreWaitGroup(t *testing.T) {
	wg := NewSemaphoreWaitGroup()
	
	wg.Add(2)
	
	done1 := make(chan bool)
	done2 := make(chan bool)
	
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		done1 <- true
	}()
	
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		done2 <- true
	}()
	
	<-done1
	<-done2
	
	wg.Wait()
}

func TestSemaphoreWaitGroupMultipleWait(t *testing.T) {
	wg := NewSemaphoreWaitGroup()
	
	wg.Add(1)
	
	waitDone1 := make(chan bool)
	waitDone2 := make(chan bool)
	
	go func() {
		wg.Wait()
		waitDone1 <- true
	}()
	
	go func() {
		wg.Wait()
		waitDone2 <- true
	}()
	
	time.Sleep(10 * time.Millisecond)
	
	wg.Done()
	
	<-waitDone1
	<-waitDone2
}
