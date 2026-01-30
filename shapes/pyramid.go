package shapes

import "sketcher/core"

type Pyramid struct {
	BaseShape
	baseLength, baseWidth, height float64
}

func NewPyramid(name string, baseLength, baseWidth, height float64) *Pyramid {
	p := &Pyramid{
		BaseShape:  NewBaseShape("Pyramid", name),
		baseLength: baseLength,
		baseWidth:  baseWidth,
		height:     height,
	}
	p.build()
	return p
}

func (p *Pyramid) build() {
	x, y, z := 0.0, 0.0, 0.0

	halfL := p.baseLength / 2.0
	halfW := p.baseWidth / 2.0

	p0 := p.mesh.AddPoint(core.Point{X: x + halfL, Y: y + halfW, Z: z}) // b1
	p1 := p.mesh.AddPoint(core.Point{X: x + halfL, Y: y - halfW, Z: z}) // b2
	p2 := p.mesh.AddPoint(core.Point{X: x - halfL, Y: y - halfW, Z: z}) // b3
	p3 := p.mesh.AddPoint(core.Point{X: x - halfL, Y: y + halfW, Z: z}) // b4

	// base
	p.mesh.AddTriangle(p0, p2, p3)
	p.mesh.AddTriangle(p2, p0, p1)

	// apex
	apex := p.mesh.AddPoint(core.Point{X: x, Y: y, Z: z + p.height})

	// sides
	p.mesh.AddTriangle(p1, p0, apex)
	p.mesh.AddTriangle(p2, p1, apex)
	p.mesh.AddTriangle(p3, p2, apex)
	p.mesh.AddTriangle(p0, p3, apex)
}
