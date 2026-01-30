package shapes

import "sketcher/core"

type Shape interface {
	Type() string
	Name() string
	Triangulation() *core.Triangulation
}
