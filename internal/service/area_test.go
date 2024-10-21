package service_test

import (
	"testing"

	"github.com/kunalmemane9150/AreaCalculator/internal/model"
	"github.com/kunalmemane9150/AreaCalculator/internal/service"
)

// CalculateArea() Tests
func TestCalculate(t *testing.T) {

	var tests = []struct {
		name              string
		shape             model.IShape
		expectedArea      string
		expectedPerimeter string
	}{
		{name: "Circle with radius 10", shape: &model.Circle{Radius: 10}, expectedArea: "314.159", expectedPerimeter: "62.832"},
		{name: "Circle with radius 15", shape: &model.Circle{Radius: 15}, expectedArea: "706.858", expectedPerimeter: "94.248"},
		{name: "Square with side 20", shape: &model.Square{Side: 20}, expectedArea: "400.000", expectedPerimeter: "80.000"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			area, perimeter := service.Calculate(tc.shape)

			if area != tc.expectedArea {
				t.Errorf("Expected Area %v, Got %v", tc.expectedArea, area)
			}
			if perimeter != tc.expectedPerimeter {
				t.Errorf("Expected Perimeter %v, Got %v", tc.expectedPerimeter, perimeter)
			}
		})
	}
}

// CalculateArea() Benchmark
func BenchmarkCalculateArea(b *testing.B) {

	shape := &model.Circle{Radius: 10}

	for i := 0; i < b.N; i++ {
		service.Calculate(shape)
	}
}
