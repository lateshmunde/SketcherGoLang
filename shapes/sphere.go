package shapes

import (
	"math"
	"sketcher/core"
	"sketcher/mathx"
)

type Sphere struct {
	BaseShape
	radius float64
}

func NewSphere(name string, radius float64) *Sphere {
	s := &Sphere{
		BaseShape: NewBaseShape("Sphere", name),
		radius:    radius,
	}
	s.build()
	return s
}

func (s *Sphere) build() {
	x, y, z := 0.0, 0.0, 0.0
	_ = x
	_ = y
	_ = z // center unused but preserved

	stacks := 36
	number := 72

	for i := 0; i < stacks; i++ {
		iLat1 := mathx.PI * (-0.5 + float64(i)/float64(stacks))
		iLat2 := mathx.PI * (-0.5 + float64(i+1)/float64(stacks))

		z1 := s.radius * math.Sin(iLat1)
		r1 := s.radius * math.Cos(iLat1)

		z2 := s.radius * math.Sin(iLat2)
		r2 := s.radius * math.Cos(iLat2)

		for j := 0; j < number; j++ {
			jLat1 := 2 * mathx.PI * float64(j) / float64(number)
			jLat2 := 2 * mathx.PI * float64(j+1) / float64(number)

			// First ring
			idx1 := s.mesh.AddPoint(core.Point{X: r1 * math.Cos(jLat1), Y: r1 * math.Sin(jLat1), Z: z1})
			idx2 := s.mesh.AddPoint(core.Point{X: r1 * math.Cos(jLat2), Y: r1 * math.Sin(jLat2), Z: z1})

			// Second ring
			idx3 := s.mesh.AddPoint(core.Point{X: r2 * math.Cos(jLat1), Y: r2 * math.Sin(jLat1), Z: z2})
			idx4 := s.mesh.AddPoint(core.Point{X: r2 * math.Cos(jLat2), Y: r2 * math.Sin(jLat2), Z: z2})

			// Same triangles
			s.mesh.AddTriangle(idx1, idx2, idx3)
			s.mesh.AddTriangle(idx2, idx4, idx3)
		}
	}
}
