package main

import (
	"log"

	"github.com/kunalmemane9150/AreaCalculator/internal/area"
	"github.com/kunalmemane9150/AreaCalculator/internal/input"
	"github.com/kunalmemane9150/AreaCalculator/internal/validator"
)

func main() {

	shape := input.GetInput()

	shape, err := validator.Validator(shape)
	if err != nil {
		log.Fatal(err)
	}

	area.CalculateArea(shape)
}
