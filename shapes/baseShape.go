package shapes

import "sketcher/core"

type BaseShape struct {
	typ  string
	name string
	mesh *core.Triangulation
}

func NewBaseShape(typ, name string) BaseShape {
	return BaseShape{
		typ:  typ,
		name: name,
		mesh: core.NewTriangulation(),
	}
}

func (b *BaseShape) Type() string {
	return b.typ
}

func (b *BaseShape) Name() string {
	return b.name
}

func (b *BaseShape) Triangulation() *core.Triangulation {
	return b.mesh
}
