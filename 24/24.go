package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

// Структура Point представляет точку с координатами x и y
type Point struct {
	x, y float64
}

// Конструктор NewPoint создает новый экземпляр точки с заданными координатами
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод DistanceTo вычисляет расстояние между текущей точкой и другой точкой
func (p *Point) distanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Задаём две точки:
	point1 := NewPoint(3.4, 4.6)
	point2 := NewPoint(-5.0, 14.2)

	// Вычисляем расстояние между точками
	distance := point1.distanceTo(point2)

	fmt.Printf("Расстояние между точкой 1 и точкой 2: %.4f\n", distance)
}
