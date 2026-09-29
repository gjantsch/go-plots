package main

import (
	"image/color"
	"time"

	graph_v1 "github.com/gjantsch/graph/cmd/09/graph"
	"github.com/gjantsch/graph/cmd/09/models"
	"github.com/zserge/fenster"
)

const (
	Width     = 500
	Height    = 500
	XMin      = 0
	YMin      = 0
	XMax      = Width - 1
	YMax      = Height - 1
	Amplitude = 100
)

func main() {

	// Open the tiniest window possible
	w, err := fenster.New(Width, Height, "Dots Canvas")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// Build the canvas object using the fenster window
	canva := graph_v1.NewCanva(w, Width, Height)

	cubeDefinition := []models.Face{
		models.NewFace(
			models.NewCoordinate(-0.5, 0.5, 0.5),
			models.NewCoordinate(0.5, 0.5, 0.5),
			models.NewCoordinate(-0.5, -0.5, 0.5),
		),
		models.NewFace(
			models.NewCoordinate(0.5, 0.5, 0.5),
			models.NewCoordinate(0.5, -0.5, 0.5),
			models.NewCoordinate(-0.5, -0.5, 0.5),
		),
		// left face
		models.NewFace(
			models.NewCoordinate(-0.5, 0.5, 0.5),
			models.NewCoordinate(-0.5, 0.5, -0.5),
			models.NewCoordinate(-0.5, -0.5, -0.5),
		),
		models.NewFace(
			models.NewCoordinate(-0.5, 0.5, 0.5),
			models.NewCoordinate(-0.5, -0.5, -0.5),
			models.NewCoordinate(-0.5, -0.5, 0.5),
		),
		// rear face
		models.NewFace(
			models.NewCoordinate(-0.5, 0.5, -0.5),
			models.NewCoordinate(0.5, 0.5, -0.5),
			models.NewCoordinate(-0.5, -0.5, -0.5),
		),
		models.NewFace(
			models.NewCoordinate(0.5, 0.5, 0.5),
			models.NewCoordinate(-0.5, -0.5, -0.5),
			models.NewCoordinate(0.5, -0.5, -0.5),
		),
		// right face
		models.NewFace(
			models.NewCoordinate(0.5, 0.5, -0.5),
			models.NewCoordinate(0.5, 0.5, 0.5),
			models.NewCoordinate(0.5, -0.5, -0.5),
		),
		models.NewFace(
			models.NewCoordinate(0.5, 0.5, -0.5),
			models.NewCoordinate(0.5, -0.5, -0.5),
			models.NewCoordinate(0.5, -0.5, 0.5),
		),
	}

	cube := models.NewObject(cubeDefinition, 1, 0)

	shader := graph_v1.NewShader(1.0)

	// Main polling frame loop (60 FPS target)
	dt := 0.016 // Approximate time per frame at 60 FPS
	for w.Loop(time.Second / 60) {
		dt += 0.016 // Increment the time delta
		// Clear screen to black every frame
		canva.Clear(color.RGBA{0, 0, 0, 255})
		shader.Distance += 0.5 * dt
		for _, face := range cube.Faces[0:2] {

			face.V1 = shader.Vertex(face.V1)
			face.V2 = shader.Vertex(face.V2)
			face.V3 = shader.Vertex(face.V3)

			canva.Rasterize(face)
		}
	}
}
