package shapes

import "sketcher/core"

type Cuboid struct {
	BaseShape
	l, w, h float64
}

func NewCuboid(name string, l, w, h float64) *Cuboid {
	c := &Cuboid{
		BaseShape: NewBaseShape("Cuboid", name),
		l:         l, w: w, h: h}
	c.build()
	return c
}

func (c *Cuboid) build() {
	x, y, z := 0.0, 0.0, 0.0

	p0 := c.mesh.AddPoint(core.NewPoint(x, y, z))
	p1 := c.mesh.AddPoint(core.NewPoint(x+c.l, y, z))
	p2 := c.mesh.AddPoint(core.NewPoint(x+c.l, y+c.w, z))
	p3 := c.mesh.AddPoint(core.NewPoint(x, y+c.w, z))

	p4 := c.mesh.AddPoint(core.NewPoint(x, y, z+c.h))
	p5 := c.mesh.AddPoint(core.NewPoint(x+c.l, y, z+c.h))
	p6 := c.mesh.AddPoint(core.NewPoint(x+c.l, y+c.w, z+c.h))
	p7 := c.mesh.AddPoint(core.NewPoint(x, y+c.w, z+c.h))

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
