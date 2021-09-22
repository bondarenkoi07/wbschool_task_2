package pattern

import (
	"fmt"
	"math"
)

type Circle struct {
	r int8
}

type Quad struct {
	a int8
}

type Triangle struct {
	a, b, c int8
}

type MathVisitor struct {
}

func (m MathVisitor) SquareCircle(c Circle) float64 {
	return math.Pow(float64(c.r), 2) * math.Pi
}

func (m MathVisitor) SquareTriangle(c Triangle) float64 {
	p := (c.a + c.b + c.c) / 2
	return math.Sqrt(float64(p * (p - c.a) * (p - c.b) * (p - c.c)))
}

func (m MathVisitor) SquareQuad(c Quad) float64 {
	return float64(c.a * c.a)
}

func init() {
	var elements = []interface{}{
		Circle{r: 10},
		Triangle{a: 3, b: 4, c: 5},
		Quad{a: 5},
	}

	var visitor MathVisitor

	for _, elem := range elements {
		switch elem.(type) {
		case Quad:
			fmt.Printf("square Quad: %f\n", visitor.SquareQuad(elem.(Quad)))
		case Triangle:
			fmt.Printf("square Triangle: %f\n", visitor.SquareTriangle(elem.(Triangle)))
		case Circle:
			fmt.Printf("square Circle: %f\n", visitor.SquareCircle(elem.(Circle)))
		}
	}
}
