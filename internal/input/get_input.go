package input

import (
	"fmt"

	"github.com/kunalmemane9150/AreaCalculator/internal/model"
)

func GetInput() model.IShape {

	var side, length, breadth, radius, sideA, sideB, sideC float64
	var shape string

	fmt.Print("Enter Shape: ")
	fmt.Scan(&shape)

	switch shape {
	case "Square":
		fmt.Print("Enter Side of Square: ")
		fmt.Scan(&side)

		return &model.Square{
			Side: side,
		}

	case "Circle":
		fmt.Print("Enter Radius of Circle: ")
		fmt.Scan(&radius)

		return &model.Circle{
			Radius: radius,
		}

	case "Rectangle":
		fmt.Print("Enter Length of Rectangle: ")
		fmt.Scan(&length)
		fmt.Print("Enter Breadth of Rectangle: ")
		fmt.Scan(&breadth)

		return &model.Rectangle{
			Length:  length,
			Breadth: breadth,
		}

	case "Triangle":
		fmt.Print("Enter Side A of Triangle: ")
		fmt.Scan(&sideA)
		fmt.Print("Enter Side B of Triangle: ")
		fmt.Scan(&sideB)
		fmt.Print("Enter Side C of Triangle: ")
		fmt.Scan(&sideC)

		return &model.Triangle{
			SideA: sideA,
			SideB: sideB,
			SideC: sideC,
		}

	default:
		return nil
	}

}
