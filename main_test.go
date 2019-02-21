package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{6, 8},
		{-10, -8},
		{0, 2},
		{100, 102},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func benchmarkCalculate(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(input)
	}
}

func BenchmarkCalculate100(b *testing.B) {
	benchmarkCalculate(100, b)
}
func BenchmarkCalculateNegative100(b *testing.B) {
	benchmarkCalculate(-100, b)
}
func BenchmarkCalculateNegative1(b *testing.B) {
	benchmarkCalculate(-1, b)
}
