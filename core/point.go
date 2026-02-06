package core

import "math"

type Point struct {
	X, Y, Z float64
}

func NewPoint(x, y, z float64) Point {
	return Point{x, y, z}
}

func (p1 Point) Distance(p2 Point) float64 {
	distSqure := (p1.X-p2.X)*(p1.X-p2.X) +
		(p1.Y-p2.Y)*(p1.Y-p2.Y) +
		(p1.Z-p2.Z)*(p1.Z-p2.Z)

	return math.Sqrt(distSqure)
}

func (p1 Point) Subtraction(p2 Point) Point {
	return Point{(p1.X - p2.X), (p1.Y - p2.Y), (p1.Z - p2.Z)}
}

func (p1 Point) IsEqualPoint(p2 Point) bool {
	if math.Abs(p1.X-p2.X) < 10e-6 && math.Abs(p1.Y-p2.X) < 10e-6 && math.Abs(p1.Z-p2.Z) < 10e-6 {
		return true
	}
	return false
}
