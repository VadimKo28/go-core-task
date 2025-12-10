package main

import (
	"testing"
)

func TestGenRandomIntWithChan(t *testing.T) {
	tests := []struct {
		name string 
		input chan int
	}{
		{
			name: "генерация числа в диапазоне 0-999",
			input: make(chan int),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go GenRandomIntWithChan(tt.input)
			
			<-tt.input
		})
	}
}

