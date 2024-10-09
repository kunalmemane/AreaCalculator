package area

import (
	"fmt"

	"github.com/kunalmemane9150/AreaCalculator/internal/model"
)

func CalculateArea(shapes model.IShape) {
	areaChan := make(chan float64)
	pChan := make(chan float64)

	go shapes.GetArea(areaChan)
	go shapes.GetPerimeter(pChan)

	shapeArea := <-areaChan
	shapePerimeter := <-pChan

	fmt.Printf("Area: %.3f\n", shapeArea)
	fmt.Printf("Perimeter: %.3f\n", shapePerimeter)
}
