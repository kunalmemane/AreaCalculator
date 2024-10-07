package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/kunalmemane9150/AreaCalculator/area"
)

func main() {

	shape, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	calculateArea(shape)
}

func getInput() (area.IShape, error) {

	var side, length, breadth, radius, sideA, sideB, sideC float64
	var shape string

	fmt.Print("Enter Shape: ")
	fmt.Scan(&shape)

	switch shape {
	case "Square":
		fmt.Print("Enter Side of Square: ")
		fmt.Scan(&side)

		if side > 0 {
			return &area.Square{
				Side: side,
			}, nil
		} else {
			return nil, errors.New("side can not be 0")
		}

	case "Circle":
		fmt.Print("Enter Radius of Circle: ")
		fmt.Scan(&radius)

		if radius > 0 {
			return &area.Circle{
				Radius: radius,
			}, nil
		} else {
			return nil, errors.New("radius can not be 0")
		}

	case "Rectangle":
		fmt.Print("Enter Length of Rectangle: ")
		fmt.Scan(&length)
		fmt.Print("Enter Breadth of Rectangle: ")
		fmt.Scan(&breadth)

		if length > 0 && breadth > 0 {
			return &area.Rectangle{
				Length:  length,
				Breadth: breadth,
			}, nil
		} else {
			return nil, errors.New("length or breadth can not be 0")
		}

	case "Triangle":
		fmt.Print("Enter Side A of Triangle: ")
		fmt.Scan(&sideA)
		fmt.Print("Enter Side B of Triangle: ")
		fmt.Scan(&sideB)
		fmt.Print("Enter Side C of Triangle: ")
		fmt.Scan(&sideC)

		if sideA > 0 && sideB > 0 && sideC > 0 {
			return &area.Triangle{
				SideA: sideA,
				SideB: sideB,
				SideC: sideC,
			}, nil
		} else {
			return nil, errors.New("sides of triangle can not be 0")
		}

	default:
		return nil, errors.New("invalid shape")
	}

}

func calculateArea(shapes area.IShape) {
	areaChan := make(chan float64)
	pChan := make(chan float64)

	go shapes.GetArea(areaChan)
	go shapes.GetPerimeter(pChan)

	shapeArea := <-areaChan
	shapePerimeter := <-pChan

	fmt.Printf("Area: %.3f\n", shapeArea)
	fmt.Printf("Perimeter: %.3f\n", shapePerimeter)
}
