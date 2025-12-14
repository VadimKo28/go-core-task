package main

import (
	"sync"
	"testing"
	"time"
)

func TestNewWaitGroup(t *testing.T) {
	wg := NewWaitGroup()
	if wg == nil {
		t.Fatal("NewWaitGroup() returned nil")
	}
	if wg.counter != 0 {
		t.Errorf("Expected counter to be 0, got %d", wg.counter)
	}
	if wg.sem == nil {
		t.Fatal("sem channel is nil")
	}
}

func TestWaitGroup_Add(t *testing.T) {
	wg := NewWaitGroup()
	
	wg.Add(1)
	if wg.counter != 1 {
		t.Errorf("Expected counter to be 1, got %d", wg.counter)
	}
}

func TestWaitGroup_Add_NegativeCounterPanic(t *testing.T) {
	wg := NewWaitGroup()
	
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected panic on negative counter, but no panic occurred")
		}
	}()
	
	wg.Add(1)
	wg.Add(-2)
}

func TestWaitGroup_Done(t *testing.T) {
	wg := NewWaitGroup()
	
	wg.Add(3)
	wg.Done()
	if wg.counter != 2 {
		t.Errorf("Expected counter to be 2, got %d", wg.counter)
	}
	
	wg.Done()
	if wg.counter != 1 {
		t.Errorf("Expected counter to be 1, got %d", wg.counter)
	}
}

func TestWaitGroup_Wait_ImmediateReturn(t *testing.T) {
	wg := NewWaitGroup()
	
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()
	
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Wait() should return immediately when counter is 0")
	}
}

func TestWaitGroup_Wait_BlocksUntilZero(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)
	
	waitDone := make(chan bool)
	go func() {
		wg.Wait()
		waitDone <- true
	}()
	
	select {
	case <-waitDone:
		t.Fatal("Wait() should block when counter > 0")
	case <-time.After(100 * time.Millisecond):
	}
	
	wg.Done()
	
	select {
	case <-waitDone:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Wait() should return after counter reaches 0")
	}
}

func TestWaitGroup_Add_ResetsSemaphore(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)
	
	waitDone1 := make(chan bool)
	go func() {
		wg.Wait()
		waitDone1 <- true
	}()
	
	wg.Done()
	
	select {
	case <-waitDone1:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("First Wait() should complete")
	}
	
	wg.Add(1)
	waitDone2 := make(chan bool)
	go func() {
		wg.Wait()
		waitDone2 <- true
	}()
	
	select {
	case <-waitDone2:
		t.Fatal("Second Wait() should block after Add(1)")
	case <-time.After(100 * time.Millisecond):
	}
	
	wg.Done()
	
	select {
	case <-waitDone2:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Second Wait() should complete after Done()")
	}
}

func TestWaitGroup_ConcurrentOperations(t *testing.T) {
	wg := NewWaitGroup()
	const numGoroutines = 100
	
	wg.Add(numGoroutines)
	
	var wgSync sync.WaitGroup
	wgSync.Add(numGoroutines)
	
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wgSync.Done()
			wg.Done()
		}()
	}
	
	wg.Wait()
	wgSync.Wait()
	
	if wg.counter != 0 {
		t.Errorf("Expected counter to be 0 after all operations, got %d", wg.counter)
	}
}

func TestWaitGroup_MultipleWait(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)
	
	const numWaiters = 10
	waitDone := make(chan bool, numWaiters)
	
	for i := 0; i < numWaiters; i++ {
		go func() {
			wg.Wait()
			waitDone <- true
		}()
	}
	
	time.Sleep(50 * time.Millisecond)
	
	wg.Done()
	
	timeout := time.After(1 * time.Second)
	for i := 0; i < numWaiters; i++ {
		select {
		case <-waitDone:
		case <-timeout:
			t.Fatalf("Only %d waiters completed, expected %d", i, numWaiters)
		}
	}
}

func TestWaitGroup_AddZero(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(0)
	
	if wg.counter != 0 {
		t.Errorf("Expected counter to be 0, got %d", wg.counter)
	}
	
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()
	
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Wait() should return immediately after Add(0)")
	}
}

func TestWaitGroup_AddAfterWait(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)
	
	waitDone1 := make(chan bool)
	go func() {
		wg.Wait()
		waitDone1 <- true
	}()
	
	wg.Done()
	
	select {
	case <-waitDone1:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("First Wait() should complete")
	}
	
	wg.Add(1)
	waitDone2 := make(chan bool)
	go func() {
		wg.Wait()
		waitDone2 <- true
	}()
	
	select {
	case <-waitDone2:
		t.Fatal("Second Wait() should block")
	case <-time.After(100 * time.Millisecond):
	}
	
	wg.Done()
	
	select {
	case <-waitDone2:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Second Wait() should complete")
	}
}

func TestWaitGroup_MultipleAdd(t *testing.T) {
	wg := NewWaitGroup()
	
	wg.Add(5)
	if wg.counter != 5 {
		t.Errorf("Expected counter to be 5, got %d", wg.counter)
	}
	
	wg.Add(-3)
	if wg.counter != 2 {
		t.Errorf("Expected counter to be 2, got %d", wg.counter)
	}
	
	wg.Add(10)
	if wg.counter != 12 {
		t.Errorf("Expected counter to be 12, got %d", wg.counter)
	}
}

func TestWaitGroup_ConcurrentAdd(t *testing.T) {
	wg := NewWaitGroup()
	const numOps = 1000
	
	var wgSync sync.WaitGroup
	wgSync.Add(numOps)
	
	for i := 0; i < numOps; i++ {
		go func() {
			defer wgSync.Done()
			wg.Add(1)
		}()
	}
	
	wgSync.Wait()
	
	if wg.counter != numOps {
		t.Errorf("Expected counter to be %d after concurrent Add(1) operations, got %d", numOps, wg.counter)
	}
	
	wgSync.Add(numOps)
	for i := 0; i < numOps; i++ {
		go func() {
			defer wgSync.Done()
			wg.Add(-1)
		}()
	}
	
	wgSync.Wait()
	
	if wg.counter != 0 {
		t.Errorf("Expected counter to be 0 after concurrent Add(-1) operations, got %d", wg.counter)
	}
}
