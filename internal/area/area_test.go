package area_test

import (
	"io"
	"os"
	"testing"

	"github.com/kunalmemane9150/AreaCalculator/internal/area"
	"github.com/kunalmemane9150/AreaCalculator/internal/model"
)

// CalculateArea() Tests
func TestCalculateArea(t *testing.T) {

	var tests = []struct {
		name     string
		shape    model.IShape
		expected string
	}{
		{name: "Circle with radius 10", shape: &model.Circle{Radius: 10}, expected: "Area: 314.159\nPerimeter: 62.832\n"},
		{name: "Circle with radius 15", shape: &model.Circle{Radius: 15}, expected: "Area: 706.858\nPerimeter: 94.248\n"},
		{name: "Square with side 20", shape: &model.Square{Side: 20}, expected: "Area: 400.000\nPerimeter: 80.000\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			stdout := os.Stdout

			r, w, _ := os.Pipe()
			os.Stdout = w

			area.CalculateArea(tc.shape)

			w.Close()

			output, err := io.ReadAll(r)
			if err != nil {
				t.Fatal(err)
			}
			os.Stdout = stdout

			if string(output) != tc.expected {
				t.Errorf("Expected %v, Got %v", tc.expected, output)
			}
		})
	}
}

// CalculateArea() Benchmark
func BenchmarkCalculateArea(b *testing.B) {

	shape := &model.Circle{Radius: 10}

	for i := 0; i < b.N; i++ {
		area.CalculateArea(shape)
	}
}
