package shapes

import "fmt"

func CreateShape(typ, name string, params ...float64) (Shape, error) {
	switch typ {
	case "Cube":
		if len(params) != 1 {
			return nil, fmt.Errorf("cube requires 1 param: side")
		}
		return NewCube(name, params[0]), nil

	case "Cuboid":
		if len(params) != 3 {
			return nil, fmt.Errorf("cuboid requires 3 params: length, width, height")
		}
		return NewCuboid(name, params[0], params[1], params[2]), nil

	case "Sphere":
		if len(params) != 1 {
			return nil, fmt.Errorf("sphere requires 1 param: radius")
		}
		return NewSphere(name, params[0]), nil

	case "Cylinder":
		if len(params) != 2 {
			return nil, fmt.Errorf("cylinder requires 2 params: radius, height")
		}
		return NewCylinder(name, params[0], params[1]), nil

	case "Cone":
		if len(params) != 2 {
			return nil, fmt.Errorf("cone requires 2 params: radius, height")
		}
		return NewCone(name, params[0], params[1]), nil

	case "Pyramid":
		if len(params) != 3 {
			return nil, fmt.Errorf("pyramid requires 3 params: length, width, height")
		}
		return NewPyramid(name, params[0], params[1], params[2]), nil
	}

	return nil, fmt.Errorf("unknown shape type: %s", typ)
}
