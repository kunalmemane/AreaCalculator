package service

import (
	"fmt"

	"github.com/kunalmemane9150/AreaCalculator/internal/model"
)

func Calculate(shapes model.IShape) (string, string) {
	areaChan := make(chan float64)
	pChan := make(chan float64)

	go shapes.GetArea(areaChan)
	go shapes.GetPerimeter(pChan)

	shapeArea := <-areaChan
	shapePerimeter := <-pChan

	area := fmt.Sprintf("%.3f", shapeArea)
	perimeter := fmt.Sprintf("%.3f", shapePerimeter)

	return area, perimeter
}
