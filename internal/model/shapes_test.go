package model_test

import (
	"testing"

	"github.com/kunalmemane9150/AreaCalculator/internal/model"
)

// GetArea() Tests
func TestGetArea(t *testing.T) {
	tests := []struct {
		name     string
		shape    model.IShape
		expected float64
	}{
		{name: "Circle Area", shape: &model.Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Square Area", shape: &model.Square{Side: 10}, expected: 100.000},
		{name: "Rectangle Area", shape: &model.Rectangle{Length: 15, Breadth: 20}, expected: 300.000},
		{name: "Triangle Area", shape: &model.Triangle{SideA: 17, SideB: 25, SideC: 20}, expected: 169.24538398432023},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			areaChan := make(chan float64)

			go tc.shape.GetArea(areaChan)

			area := <-areaChan

			if area != tc.expected {
				t.Errorf("Expected %v, Got %v", tc.expected, area)
			}
		})

	}
}

// GetArea() Benchmark
func BenchmarkGetArea(b *testing.B) {

	shape := &model.Circle{Radius: 10}
	areaChan := make(chan float64)

	for i := 0; i < b.N; i++ {
		go shape.GetArea(areaChan)
		<-areaChan
	}
}

// GetPerimeter Tests
func TestGetPerimeter(t *testing.T) {
	tests := []struct {
		name     string
		shape    model.IShape
		expected float64
	}{
		{name: "Circle Perimeter", shape: &model.Circle{Radius: 10}, expected: 62.83185307179586},
		{name: "Square Perimeter", shape: &model.Square{Side: 10}, expected: 40.000},
		{name: "Rectangle Perimeter", shape: &model.Rectangle{Length: 10, Breadth: 20}, expected: 60.000},
		{name: "Triangle Perimeter", shape: &model.Triangle{SideA: 30, SideB: 20, SideC: 50}, expected: 100.000},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			pChan := make(chan float64)

			go tc.shape.GetPerimeter(pChan)

			perimeter := <-pChan

			if perimeter != tc.expected {
				t.Errorf("Expected %v, Got %v", tc.expected, perimeter)
			}
		})

	}
}

// GetPerimeter Benchmark
func BenchmarkGetPerimeter(b *testing.B) {

	shape := &model.Circle{Radius: 10}
	pChan := make(chan float64)

	for i := 0; i < b.N; i++ {
		go shape.GetPerimeter(pChan)
		<-pChan
	}
}

// Test constructors ---
func TestNewSquare(t *testing.T) {
	tests := []struct {
		name         string
		inputSide    float64
		expectedSide float64
	}{{name: "Square side value", inputSide: 10, expectedSide: 10}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			sq := model.NewSquare(tc.inputSide)

			if sq.Side != tc.expectedSide {
				t.Errorf("Expected %v, got %v", tc.expectedSide, sq.Side)
			}
		})
	}
}
func TestNewCircle(t *testing.T) {
	tests := []struct {
		name           string
		inputRadius    float64
		expectedRadius float64
	}{{name: "Circle Radius value", inputRadius: 20, expectedRadius: 20}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			cir := model.NewCircle(tc.inputRadius)

			if cir.Radius != tc.expectedRadius {
				t.Errorf("Expected %v, got %v", tc.expectedRadius, cir.Radius)
			}
		})
	}
}
func TestNewRectangle(t *testing.T) {
	tests := []struct {
		name            string
		inputLength     float64
		inputBreadth    float64
		expectedLength  float64
		expectedBreadth float64
	}{{name: "Rectangle Length and Breadth value", inputLength: 20, inputBreadth: 30, expectedLength: 20, expectedBreadth: 30}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			rect := model.NewRectangle(tc.inputLength, tc.inputBreadth)

			if rect.Length != tc.expectedLength {
				t.Errorf("Expected %v, got %v", tc.expectedLength, rect.Length)
			}
			if rect.Breadth != tc.expectedBreadth {
				t.Errorf("Expected %v, got %v", tc.expectedBreadth, rect.Breadth)
			}
		})
	}
}
func TestNewTriangle(t *testing.T) {
	tests := []struct {
		name          string
		inputSideA    float64
		inputSideB    float64
		inputSideC    float64
		expectedSideA float64
		expectedSideB float64
		expectedSideC float64
	}{{name: "Circle Radius value", inputSideA: 20, inputSideB: 25, inputSideC: 30, expectedSideA: 20, expectedSideB: 25, expectedSideC: 30}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			tri := model.NewTriangle(tc.inputSideA, tc.inputSideB, tc.inputSideC)

			if tri.SideA != tc.expectedSideA {
				t.Errorf("Expected %v, got %v", tc.expectedSideA, tri.SideA)
			}
			if tri.SideB != tc.expectedSideB {
				t.Errorf("Expected %v, got %v", tc.expectedSideB, tri.SideB)
			}
			if tri.SideC != tc.expectedSideC {
				t.Errorf("Expected %v, got %v", tc.expectedSideC, tri.SideC)
			}
		})
	}
}
