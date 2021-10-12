package pattern

import (
	"fmt"
	"math"
)

// Посетитель — это поведенческий паттерн проектирования,
//который позволяет добавлять в программу новые операции,
//не изменяя классы объектов, над которыми эти операции могут выполняться.

// Плюсы:
//Упрощает добавление операций, работающих со сложными структурами объектов.
//Объединяет родственные операции в одном классе.
//Посетитель может накапливать состояние при обходе структуры элементов.
// Минусы:
//Паттерн не оправдан, если иерархия элементов часто меняется.
//Может привести к нарушению инкапсуляции элементов.

type Circle struct {
	r int8
}

type Quad struct {
	a int8
}

type Triangle struct {
	a, b, c int8
}

// MathVisitor находит площадь фигуры, в зависимости от ее типа
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
