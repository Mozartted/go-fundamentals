package structsmethodsinterfaces

import "math"

type Shape interface {
	// Perimeter() float64
	Area() float64
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.base * t.height)
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return (r.width * r.height)
}

type Circle struct {
	radius float64
}

func (r Circle) Perimeter() float64 {
	return 2 * r.radius * math.Pi
}

func (r Circle) Area() float64 {
	return math.Pi * math.Pow(r.radius, 2.0)
}

func Perimeter(rec *Rectangle) float64 {
	return 2 * (rec.height + rec.width)
}

func Area(rec *Rectangle) float64 {
	return rec.height * rec.width
}
