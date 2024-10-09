package input_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/kunalmemane9150/AreaCalculator/internal/input"
	"github.com/kunalmemane9150/AreaCalculator/internal/model"
)

// GetInput() Tests
func TestGetInput(t *testing.T) {
	var tests = []struct {
		name          string
		input         string
		expectedShape model.IShape
	}{
		{name: "Circle", input: "Circle\n10\n", expectedShape: &model.Circle{}},

		{name: "Square", input: "Square\n10\n", expectedShape: &model.Square{}},

		{name: "Rectangle", input: "Rectangle\n10\n20\n", expectedShape: &model.Rectangle{}},

		{name: "Triangle", input: "Triangle\n10\n20\n30\n", expectedShape: &model.Triangle{}},

		{name: "Hexagon", input: "Hexagon\n", expectedShape: nil},
	}

	for _, tc := range tests {

		t.Run(tc.name+": "+tc.name, func(t *testing.T) {

			reader, writer, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}

			if _, err := writer.Write([]byte(tc.input)); err != nil {
				t.Fatal(err)
			}
			os.Stdin = reader

			shape := input.GetInput()

			if reflect.TypeOf(shape) != reflect.TypeOf(tc.expectedShape) {
				t.Errorf("Expected Shape %s, Got %s", tc.expectedShape, reflect.TypeOf(shape))
			}
		})
	}
}

// GetInput() Benchmark
func BenchmarkGetInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input.GetInput()
	}
}
