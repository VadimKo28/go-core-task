package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge_Basic(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

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

	merged := Merge(ch1, ch2, ch3)

	var result []int
	for v := range merged {
		result = append(result, v)
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Ints(result)
	sort.Ints(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMerge_EmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		close(ch1)
	}()
	go func() {
		close(ch2)
	}()

	merged := Merge(ch1, ch2)

	var result []int
	for v := range merged {
		result = append(result, v)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
}

