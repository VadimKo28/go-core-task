package main

import (
	"fmt"
	"sync"
	// "time"
	// "time"
)

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()

	go func() {
		defer close(ch2)
		ch2 <- 4
		ch2 <- 5
		ch2 <- 6
	}()

	go func() {
		defer close(ch3)
		ch3 <- 7
		ch3 <- 8
		ch3 <- 9
	}()

	fmt.Println("Сливаем данные из каналов в один")
	mergedCh := Merge(ch1, ch2, ch3)

	for v := range mergedCh {
		fmt.Println(v)
	}
}


func Merge(channels ...<-chan int) <-chan int {
	merged := make(chan int, 1)

  wg := sync.WaitGroup{}

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			for num := range ch {
					merged <- num
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}