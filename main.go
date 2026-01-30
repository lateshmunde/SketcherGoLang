package main

import (
	"sketcher/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// API
	r.POST("/shape", api.CreateShapeHandler)

	// Serve frontend
	r.Static("/web", "./web") // http://localhost:8080/web/viewer.html

	r.Run(":8080")
}
