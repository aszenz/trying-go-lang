package interfaces

import "math"

// Shape defines a geometric shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle defines a shape with width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle defines a shape with radius
type Circle struct {
	Radius float64
}

// Triangle defines a shape with base and altitude
type Triangle struct {
	Base     float64
	Altitude float64
}

// Perimeter returns the permiter of the rectangle
func (rect Rectangle) Perimeter() (res float64) {
	return 2 * (rect.Width + rect.Height)
}

// Area returns the area of the rectangle
func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the area of the circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area returns the area of the triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Altitude
}

// Perimeter returns the perimeter of the triangle
func (t Triangle) Perimeter() float64 {
	return 0
}
