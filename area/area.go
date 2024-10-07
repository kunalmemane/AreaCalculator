package area

import "math"

type IShape interface {
	GetArea(chan float64)
	GetPerimeter(chan float64)
}

type Square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

// Square
func (square *Square) GetArea(areaChan chan float64) {
	sqArea := square.Side * square.Side
	areaChan <- sqArea
}

func (square *Square) GetPerimeter(periChan chan float64) {
	sqPerimeter := 4 * square.Side
	periChan <- sqPerimeter
}

// Circle
func (circle *Circle) GetArea(areaChan chan float64) {
	cirArea := math.Pi * math.Pow(circle.Radius, 2)
	areaChan <- cirArea
}

func (circle *Circle) GetPerimeter(periChan chan float64) {
	cirPerimeter := 2 * math.Pi * circle.Radius
	periChan <- cirPerimeter
}

// Rectangle
func (rect *Rectangle) GetArea(areaChan chan float64) {
	rectArea := rect.Length * rect.Breadth
	areaChan <- rectArea
}

func (rect *Rectangle) GetPerimeter(periChan chan float64) {
	rectPerimeter := 2 * (rect.Length + rect.Breadth)
	periChan <- rectPerimeter
}

// Triangle
func (tri *Triangle) GetArea(areaChan chan float64) {
	s := (tri.SideA + tri.SideB + tri.SideC) / 2
	triangleArea := math.Sqrt(s * (s - tri.SideA) * (s - tri.SideB) * (s - tri.SideC))
	areaChan <- triangleArea
}

func (tri *Triangle) GetPerimeter(periChan chan float64) {
	triPerimeter := tri.SideA + tri.SideB + tri.SideC
	periChan <- triPerimeter
}
