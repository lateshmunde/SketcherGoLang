package shapes

import "sketcher/core"

type Cube struct {
	BaseShape
	side float64
}

func NewCube(name string, side float64) *Cube {
	c := &Cube{
		BaseShape: NewBaseShape("Cube", name),
		side:      side,
	}
	c.build()
	return c
}

func (c *Cube) build() {
	x, y, z := 0.0, 0.0, 0.0

	p0 := c.mesh.AddPoint(core.NewPoint(x, y, z))
	p1 := c.mesh.AddPoint(core.NewPoint(x+c.side, y, z))
	p2 := c.mesh.AddPoint(core.NewPoint(x+c.side, y+c.side, z))
	p3 := c.mesh.AddPoint(core.NewPoint(x, y+c.side, z))

	p4 := c.mesh.AddPoint(core.NewPoint(x, y, z+c.side))
	p5 := c.mesh.AddPoint(core.NewPoint(x+c.side, y, z+c.side))
	p6 := c.mesh.AddPoint(core.NewPoint(x+c.side, y+c.side, z+c.side))
	p7 := c.mesh.AddPoint(core.NewPoint(x, y+c.side, z+c.side))

	c.mesh.AddTriangle(p0, p2, p1)
	c.mesh.AddTriangle(p0, p3, p2)

	c.mesh.AddTriangle(p4, p5, p6)
	c.mesh.AddTriangle(p4, p6, p7)

	c.mesh.AddTriangle(p7, p6, p2)
	c.mesh.AddTriangle(p7, p2, p3)

	c.mesh.AddTriangle(p0, p1, p5)
	c.mesh.AddTriangle(p0, p5, p4)

	c.mesh.AddTriangle(p5, p1, p2)
	c.mesh.AddTriangle(p5, p2, p6)

	c.mesh.AddTriangle(p0, p4, p7)
	c.mesh.AddTriangle(p0, p7, p3)
}
