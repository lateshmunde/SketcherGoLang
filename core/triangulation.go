package core

import "math"

type Triangulation struct {
	points     []Point
	triangles  []Triangle
	pointIndex map[Point]int
}

func NewTriangulation() *Triangulation {
	return &Triangulation{
		pointIndex: make(map[Point]int),
	}
}

func (t *Triangulation) AddPoint(p Point) int {
	if idx, ok := t.pointIndex[p]; ok {
		return idx
	}
	idx := len(t.points)
	t.points = append(t.points, p)
	t.pointIndex[p] = idx
	return idx
}

func (t *Triangulation) AddTriangle(a, b, c int) {
	tri := Triangle{M1: a, M2: b, M3: c}
	tri.Normal = t.calculateNormal(tri)
	t.triangles = append(t.triangles, tri)
}

func (t *Triangulation) calculateNormal(tr Triangle) Point {
	u := t.points[tr.M2].Subtraction(t.points[tr.M1])
	v := t.points[tr.M3].Subtraction(t.points[tr.M1])

	nx := u.Y*v.Z - u.Z*v.Y
	ny := u.Z*v.X - u.X*v.Z
	nz := u.X*v.Y - u.Y*v.X

	len := math.Sqrt(nx*nx + ny*ny + nz*nz)

	return Point{nx / len, ny / len, nz / len}
}

func (t *Triangulation) Points() []Point {
	return t.points
}

func (t *Triangulation) Triangles() []Triangle {
	return t.triangles
}
