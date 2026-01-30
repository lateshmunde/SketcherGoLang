package api

import (
	"net/http"
	"sketcher/shapes"

	"github.com/gin-gonic/gin"
)

func CreateShapeHandler(c *gin.Context) {
	var req ShapeRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shape, err := shapes.CreateShape(req.Type, req.Name, req.Params...)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mesh := shape.Triangulation()

	// convert to data transfer objects(DTO)
	points := make([]PointDTO, 0)
	for _, p := range mesh.Points() {
		points = append(points, PointDTO{p.X, p.Y, p.Z})
	}

	tris := make([]TriangleDTO, 0)
	for _, t := range mesh.Triangles() {
		tris = append(tris, TriangleDTO{t.M1, t.M2, t.M3})
	}

	resp := ShapeResponse{
		Type:      shape.Type(),
		Name:      shape.Name(),
		Points:    points,
		Triangles: tris,
	}

	c.JSON(http.StatusOK, resp)
}
