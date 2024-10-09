package validator_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kunalmemane9150/AreaCalculator/internal/model"
	"github.com/kunalmemane9150/AreaCalculator/internal/validator"
)

// GetInput() Tests
func TestValidator(t *testing.T) {

	var tests = []struct {
		name      string
		shapeName string
		shape     model.IShape

		expectedShape model.IShape
		expectedErr   error
	}{
		{name: "Positive radius", shapeName: "Circle", shape: &model.Circle{Radius: 10}, expectedShape: &model.Circle{}, expectedErr: nil},
		{name: "Negative radius", shapeName: "Circle", shape: &model.Circle{Radius: -10}, expectedShape: nil, expectedErr: errors.New("invalid circle radius")},

		{name: "Positive side", shapeName: "Square", shape: &model.Square{Side: 20}, expectedShape: &model.Square{}, expectedErr: nil},
		{name: "Negative side", shapeName: "Square", shape: &model.Square{Side: -20}, expectedShape: nil, expectedErr: errors.New("invalid square side")},

		{name: "Positive Length & Breadth", shapeName: "Rectangle", shape: &model.Rectangle{Length: 10, Breadth: 20}, expectedShape: &model.Rectangle{}, expectedErr: nil},
		{name: "Negative Length, Positive Breadth", shape: &model.Rectangle{Length: -10, Breadth: 20}, shapeName: "Rectangle", expectedShape: nil, expectedErr: errors.New("invalid rectangle params")},
		{name: "Positive Length, Negative Breadth", shapeName: "Rectangle", shape: &model.Rectangle{Length: 10, Breadth: -20}, expectedShape: nil, expectedErr: errors.New("invalid rectangle params")},
		{name: "Both Negative Length & Breadth", shapeName: "Rectangle", shape: &model.Rectangle{Length: -20, Breadth: -30}, expectedShape: nil, expectedErr: errors.New("invalid rectangle params")},

		{name: "Negative SideA", shapeName: "Triangle", shape: &model.Triangle{SideA: -10}, expectedShape: nil, expectedErr: errors.New("invalid triangle side")},
		{name: "Negative SideB", shapeName: "Triangle", shape: &model.Triangle{SideB: -20}, expectedShape: nil, expectedErr: errors.New("invalid triangle side")},
		{name: "Negative SideC", shapeName: "Triangle", shape: &model.Triangle{SideC: -30}, expectedShape: nil, expectedErr: errors.New("invalid triangle side")},
		{name: "Positive SideA, SideB, SideC", shapeName: "Triangle", shape: &model.Triangle{SideA: 15, SideB: 20, SideC: 25}, expectedShape: &model.Triangle{}, expectedErr: nil},
		{name: "Negative SideA, SideB, SideC", shapeName: "Triangle", shape: &model.Triangle{SideA: -15, SideB: -20, SideC: -25}, expectedShape: nil, expectedErr: errors.New("invalid triangle side")},

		{name: "Invalid Shape", shapeName: "Hexagon", shape: nil, expectedShape: nil, expectedErr: errors.New("invalid shape")},
	}

	for _, tc := range tests {

		t.Run(tc.shapeName+": "+tc.name, func(t *testing.T) {

			shape, err := validator.Validator(tc.shape)
			if err != nil {
				if err.Error() != tc.expectedErr.Error() {
					t.Errorf("Expected %v, Got %v", tc.expectedErr, err)
				}
			}

			if reflect.TypeOf(shape) != reflect.TypeOf(tc.expectedShape) {
				t.Errorf("Expected Shape %s, Got %s", tc.expectedShape, reflect.TypeOf(shape))
			}
		})
	}
}

// GetInput() Benchmark
func BenchmarkValidator(b *testing.B) {
	shape := model.Circle{}
	for i := 0; i < b.N; i++ {
		validator.Validator(&shape)
	}
}
