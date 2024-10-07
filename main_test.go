package main

import (
	"os"
	"testing"
)

func TestGetInput(t *testing.T) {
	var tests = []struct {
		name       string
		shape      string
		parameters int
	}{
		{name: "Negative circle radius", shape: "Circle", parameters: -1},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			reader, writer, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}

			if _, err := writer.Write([]byte{byte(tc.parameters)}); err != nil {
				t.Fatal(err)
			}

			os.Stdin = reader

			getInput()

		})

	}

}
