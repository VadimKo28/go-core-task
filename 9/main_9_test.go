package main

import (
	"math"
	"testing"
)

func TestPipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go pipeline(in, out)

	testCases := []struct {
		input  uint8
		output float64
	}{
		{1, 1.0},
		{2, 8.0},
		{3, 27.0},
		{4, 64.0},
		{5, 125.0},
		{10, 1000.0},
		{0, 0.0},
	}

	go func() {
		defer close(in)
		for _, tc := range testCases {
			in <- tc.input
		}
	}()

	for _, tc := range testCases {
		result := <-out
		expected := tc.output
		if math.Abs(result-expected) > 0.0001 {
			t.Errorf("Для входного значения %d ожидалось %.4f, получено %.4f", tc.input, expected, result)
		}
	}
}

func TestPipelineEmptyChannel(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go pipeline(in, out)
	close(in)

	result, ok := <-out
	if ok {
		t.Errorf("Ожидалось, что канал будет закрыт, но получено значение: %f", result)
	}
}
