package main

import (
	"sort"
	"testing"
)

// helper: drain channel into slice
func sliceFromChan(ch <-chan int) []int {
	var res []int
	for v := range ch {
		res = append(res, v)
	}
	return res
}

// helper: create channels with values
func createChannelWithValues(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	return ch
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name           string
		numChannels    int
		values         [][]int
		expectedResult []int
	}{
		{
			name:           "merge 3 channels",
			numChannels:    3,
			values:         [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expectedResult: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:           "merge zero channels",
			numChannels:    0,
			values:         [][]int{},
			expectedResult: []int{},
		},
		{
			name:           "merge 4 channels",
			numChannels:    4,
			values:         [][]int{{1, 2}, {3, 4}, {5}, {6, 7, 8}},
			expectedResult: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var channels []<-chan int
			for i := 0; i < tt.numChannels; i++ {
				if i < len(tt.values) {
					channels = append(channels, createChannelWithValues(tt.values[i]...))
				}
			}

			mergedCh := Merge(channels...)

			got := sliceFromChan(mergedCh)
			sort.Ints(got)

			if len(got) != len(tt.expectedResult) {
				t.Errorf("unexpected length: got %d want %d", len(got), len(tt.expectedResult))
			}

			for i := range tt.expectedResult {
				if i < len(got) && got[i] != tt.expectedResult[i] {
					t.Errorf("mismatch at index %d: got %v want %v", i, got, tt.expectedResult)
					break
				}
			}
		})
	}
}
