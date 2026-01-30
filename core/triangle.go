package core

type Triangle struct {
	M1, M2, M3 int
	Normal     Point
}

func NewTriangle(a, b, c int, n Point) Triangle {
	return Triangle{a, b, c, n}
}
