package model

import (
	"math"
)

type IShape interface {
	GetArea(chan float64)
	GetPerimeter(chan float64)
}

// Request model
type ShapeRequest struct {
	Shape   string  `json:"shape"`
	Side    float64 `json:"side"`
	Radius  float64 `json:"radius"`
	Length  float64 `json:"length"`
	Breadth float64 `json:"breadth"`
	SideA   float64 `json:"sideA"`
	SideB   float64 `json:"sideB"`
	SideC   float64 `json:"sideC"`
}

// Shape Dimension model
type ShapeDimensions struct {
	Side    float64 `json:"side,omitempty"`
	Radius  float64 `json:"radius,omitempty"`
	Length  float64 `json:"length,omitempty"`
	Breadth float64 `json:"breadth,omitempty"`
	SideA   float64 `json:"sideA,omitempty"`
	SideB   float64 `json:"sideB,omitempty"`
	SideC   float64 `json:"sideC,omitempty"`
}

// Final shape response
type ShapeResponse struct {
	Shape      string          `json:"shape"`
	Dimensions ShapeDimensions `json:"dimensions"`
	Area       string          `json:"area"`
	Perimeter  string          `json:"perimeter"`
}

// Final API response
type APIResponse struct {
	Success   bool          `json:"success"`
	Shape     ShapeResponse `json:"data"`
	Timestamp string        `json:"timestamp"`
}

// Error response
type ErrorResponse struct {
	Success    bool   `json:"success"`
	Error      string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

// Square
type Square struct {
	Side float64
}

func NewSquare(side float64) *Square {
	return &Square{Side: side}
}

// Circle
type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

// Rectangle
type Rectangle struct {
	Length  float64
	Breadth float64
}

func NewRectangle(length float64, breadth float64) *Rectangle {
	return &Rectangle{Length: length, Breadth: breadth}
}

// Triangle
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func NewTriangle(sideA float64, sideB float64, sideC float64) *Triangle {
	return &Triangle{SideA: sideA, SideB: sideB, SideC: sideC}
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
