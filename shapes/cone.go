package shapes

import (
	"math"
	"sketcher/core"
	"sketcher/mathx"
)

type Cone struct {
	BaseShape
	radius, height float64
}

func NewCone(name string, radius, height float64) *Cone {
	c := &Cone{
		BaseShape: NewBaseShape("Cone", name),
		radius:    radius,
		height:    height,
	}
	c.build()
	return c
}

func (c *Cone) build() {
	x, y, z := 0.0, 0.0, 0.0

	// Origin and Apex
	origin := c.mesh.AddPoint(core.Point{X: x, Y: y, Z: z})
	apex := c.mesh.AddPoint(core.Point{X: x, Y: y, Z: z + c.height})

	// First base point
	var basePts []int
	basePts = append(basePts,
		c.mesh.AddPoint(core.Point{
			X: x + c.radius*math.Cos(0),
			Y: y + c.radius*math.Sin(0),
			Z: z,
		}),
	)

	number := 72
	dTheta := 2 * mathx.PI / float64(number)

	for i := 1; i <= number; i++ {
		theta := float64(i) * dTheta

		x1 := c.radius * math.Cos(theta)
		y1 := c.radius * math.Sin(theta)

		basePts = append(basePts,
			c.mesh.AddPoint(core.Point{X: x + x1, Y: y + y1, Z: z}),
		)

		// Base triangle
		c.mesh.AddTriangle(basePts[i-1], origin, basePts[i])

		// Side triangle
		c.mesh.AddTriangle(basePts[i-1], basePts[i], apex)
	}
}
