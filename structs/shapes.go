package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// burada circle Circle kısmı go'da receiver diye geçiyor
//	Circle struct'ına ait bir fonksiyon oluyor
func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}
