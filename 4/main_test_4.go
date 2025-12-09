package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "example from task",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "empty slice1",
			slice1:   []string{},
			slice2:   []string{"banana", "date"},
			expected: []string{},
		},
		{
			name:     "empty slice2",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{},
			expected: []string{"apple", "banana"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Difference() = %v, want %v", result, tt.expected)
			}
		})
	}
}
