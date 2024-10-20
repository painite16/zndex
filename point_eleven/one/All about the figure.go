package code

import (
	"math"
)

type Circle struct {
	radius float64
}
type Rectangle struct {
	width  float64
	height float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

//Создайте интерфейс Shape с методом Area, который будет возвращать площадь фигуры. Создайте структуры Circle и Rectangle, которые будут реализовывать этот интерфейс и рассчитывать площадь этих фигур.
