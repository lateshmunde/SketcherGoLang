package shapes

import (
	"math"
	"sketcher/core"
	"sketcher/mathx"
)

type Cylinder struct {
	BaseShape
	radius float64
	height float64
}

func NewCylinder(name string, radius, height float64) *Cylinder {
	cy := &Cylinder{
		BaseShape: NewBaseShape("Cylinder", name),
		radius:    radius,
		height:    height,
	}
	cy.build()
	return cy
}

func (c *Cylinder) build() {
	x, y, z := 0.0, 0.0, 0.0

	// Centers
	baseCenter := c.mesh.AddPoint(core.Point{X: x, Y: y, Z: z})
	topCenter := c.mesh.AddPoint(core.Point{X: x, Y: y, Z: z + c.height})

	var bPts []int
	var tPts []int

	// First ring point
	bPts = append(bPts, c.mesh.AddPoint(core.Point{
		X: x + c.radius*math.Cos(0),
		Y: y + c.radius*math.Sin(0),
		Z: z,
	}))
	tPts = append(tPts, c.mesh.AddPoint(core.Point{
		X: x + c.radius*math.Cos(0),
		Y: y + c.radius*math.Sin(0),
		Z: z + c.height,
	}))

	number := 72
	dTheta := 2 * mathx.PI / float64(number)

	for i := 1; i <= number; i++ {
		theta := float64(i) * dTheta

		x1 := c.radius * math.Cos(theta)
		y1 := c.radius * math.Sin(theta)

		bPts = append(bPts, c.mesh.AddPoint(core.Point{X: x + x1, Y: y + y1, Z: z}))
		tPts = append(tPts, c.mesh.AddPoint(core.Point{X: x + x1, Y: y + y1, Z: z + c.height}))
		// Same triangle order as C++
		c.mesh.AddTriangle(bPts[i], bPts[i-1], baseCenter)
		c.mesh.AddTriangle(bPts[i-1], bPts[i], tPts[i])
		c.mesh.AddTriangle(bPts[i-1], tPts[i], tPts[i-1])
		c.mesh.AddTriangle(topCenter, tPts[i-1], tPts[i])
	}
}
