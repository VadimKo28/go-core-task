package main

import (
	"reflect"
	"testing"
)

func TestExistUnion(t *testing.T) {
	tests := []struct {
		name     string
		s1       []int
		s2       []int
		wantOk   bool
		wantResult []int
	}{
		{
			name:     "есть пересечения",
			s1:       []int{65, 3, 58, 678, 64},
			s2:       []int{64, 2, 3, 43},
			wantOk:   true,
			wantResult: []int{64, 3},
		},
		{
			name:     "нет пересечений",
			s1:       []int{1, 2, 3},
			s2:       []int{4, 5, 6},
			wantOk:   true,
			wantResult: []int{},
		},
		{
			name:     "первый слайс пустой",
			s1:       []int{},
			s2:       []int{1, 2, 3},
			wantOk:   false,
			wantResult: []int{},
		},
		{
			name:     "второй слайс пустой",
			s1:       []int{1, 2, 3},
			s2:       []int{},
			wantOk:   false,
			wantResult: []int{},
		},
		{
			name:     "оба слайса пустые",
			s1:       []int{},
			s2:       []int{},
			wantOk:   false,
			wantResult: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotResult := ExistUnion(tt.s1, tt.s2)
			if gotOk != tt.wantOk {
				t.Errorf("ExistUnion() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ExistUnion() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
