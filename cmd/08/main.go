package main

import (
	"image/color"
	"time"

	graph_v1 "github.com/gjantsch/graph/pkg/graph/v1"
	models "github.com/gjantsch/graph/pkg/models/v1"
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

	// greenColor := graph_v1.NewShader(0.5, 1, 0.5, 0.1)
	blueColor := graph_v1.NewShader(1, 1, 1, 1)

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

	// Main polling frame loop (60 FPS target)
	for w.Loop(time.Second / 60) {

		// Clear screen to black every frame
		canva.Clear(color.RGBA{0, 0, 0, 255})

		for _, face := range cube.Faces[0:2] {
			v1, v2, v3 := canva.FaceToScreen(face)
			canva.Rasterize(v1, v2, v3, blueColor)
		}
	}
}
