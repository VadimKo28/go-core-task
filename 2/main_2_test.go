package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sliceExample(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "even numbers from slice",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "only odd numbers",
			input:    []int{1, 3, 5, 7},
			expected: []int{},
		},
		{
			name:     "only even numbers",
			input:    []int{2, 4, 6},
			expected: []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceExample(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sliceExample(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func Test_addElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		element  int
		expected []int
	}{
		{
			name:     "add element to slice",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			element:  99,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 99},
		},
		{
			name:     "add element to empty slice",
			input:    []int{},
			element:  42,
			expected: []int{42},
		},
		{
			name:     "add element to single element slice",
			input:    []int{1},
			element:  2,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addElements(tt.input, tt.element)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("addElements(%v, %d) = %v, expected %v", tt.input, tt.element, result, tt.expected)
			}
		})
	}
}

func Test_copySlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "copy slice",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "copy empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "copy single element slice",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := copySlice(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("copySlice(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
			if len(tt.input) > 0 && fmt.Sprintf("%p", tt.input) == fmt.Sprintf("%p", result) {
				t.Errorf("copySlice(%v) should return a different slice, but pointers are equal", tt.input)
			}
			if len(tt.input) > 0 {
				modified := append(tt.input, 100)
				if len(result) != len(tt.expected) {
					t.Errorf("copySlice(%v) result was modified when original slice changed", tt.input)
				}
				if len(modified) == len(result) && len(result) > 0 && &result[0] == &modified[0] {
					t.Errorf("copySlice(%v) result shares underlying array with modified slice", tt.input)
				}
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		index    int
		expected []int
	}{
		{
			name:     "remove middle element",
			input:    []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "remove first element",
			input:    []int{1, 2, 3, 4, 5},
			index:    0,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "remove last element",
			input:    []int{1, 2, 3, 4, 5},
			index:    4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "negative index",
			input:    []int{1, 2, 3},
			index:    -1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "index out of bounds",
			input:    []int{1, 2, 3},
			index:    10,
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty slice",
			input:    []int{},
			index:    0,
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{42},
			index:    0,
			expected: []int{},
		},
		{
			name:     "single element invalid index",
			input:    []int{42},
			index:    1,
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeElement(tt.input, tt.index)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("removeElement(%v, %d) = %v, expected %v", tt.input, tt.index, result, tt.expected)
			}
		})
	}
}
