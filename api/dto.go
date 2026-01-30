package api

type ShapeRequest struct {
	Type   string    `json:"type"`
	Name   string    `json:"name"`
	Params []float64 `json:"params"`
}

type PointDTO struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type TriangleDTO struct {
	M1 int `json:"a"`
	M2 int `json:"b"`
	M3 int `json:"c"`
}

type ShapeResponse struct {
	Type      string        `json:"type"`
	Name      string        `json:"name"`
	Points    []PointDTO    `json:"points"`
	Triangles []TriangleDTO `json:"triangles"`
}
