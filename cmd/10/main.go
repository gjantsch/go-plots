package main

import (
	"image/color"
	"time"

	graph_v1 "github.com/gjantsch/graph/cmd/10/graph"
	"github.com/gjantsch/graph/cmd/10/models"
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

	ColorRed := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	ColorGreen := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	ColorBlue := color.RGBA{R: 0, G: 0, B: 255, A: 255}
	ColorPurple := color.RGBA{R: 255, G: 0, B: 255, A: 255}
	ColorGray := color.RGBA{R: 128, G: 128, B: 128, A: 255}
	ColorYellow := color.RGBA{R: 255, G: 255, B: 0, A: 255}

	// Open the tiniest window possible
	w, err := fenster.New(Width, Height, "Dots Canvas")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// Build the canvas object using the fenster window
	canva := graph_v1.NewCanva(w, Width, Height)

	va := models.NewCoordinate(0.5, 0.5, 0.5)
	vb := models.NewCoordinate(0.5, -0.5, 0.5)
	vc := models.NewCoordinate(-0.5, -0.5, 0.5)
	vd := models.NewCoordinate(-0.5, 0.5, 0.5)
	ve := models.NewCoordinate(-0.5, 0.5, -0.5)
	vf := models.NewCoordinate(-0.5, -0.5, -0.5)
	vg := models.NewCoordinate(0.5, -0.5, -0.5)
	vh := models.NewCoordinate(0.5, 0.5, -0.5)

	cubeDefinition := []models.Face{
		// right face
		models.NewFace(vg, vh, va, ColorPurple),
		models.NewFace(vg, va, vb, ColorPurple),
		// top face
		models.NewFace(vd, vh, va, ColorGray),
		models.NewFace(vd, vh, ve, ColorGray),
		// front face
		models.NewFace(va, vb, vc, ColorRed),
		models.NewFace(va, vc, vd, ColorRed),
		// top face
		models.NewFace(vf, vc, vb, ColorYellow),
		models.NewFace(vf, vg, vb, ColorYellow),
		// rear face
		models.NewFace(ve, vh, vg, ColorGreen),
		models.NewFace(ve, vf, vg, ColorGreen),
		// left face
		models.NewFace(vd, vc, vf, ColorBlue),
		models.NewFace(vd, vf, ve, ColorBlue),
	}

	cube := models.NewObject(cubeDefinition, 1, 0)

	// Main polling frame loop (60 FPS target)
	dt := 0.033 // Approximate time per frame at 30 FPS
	rotation := 0.0
	for w.Loop(time.Second / 30) {
		dt = 0.033 // Increment the time delta
		// Clear screen to black every frame
		canva.Clear(color.RGBA{0, 0, 0, 255})
		for j, face := range cube.Faces {
			rotation += dt * 0.05
			shader := graph_v1.NewShader(1.8, color.RGBA{R: 255, G: uint8(j * 10), B: 0, A: 255})
			shader.RotationX = 3.14 * 45 / 180
			shader.RotationY = -3.14 * 45 / 180
			shader.RotationZ = 0
			face.V1 = shader.Vertex(face.V1)
			face.V2 = shader.Vertex(face.V2)
			face.V3 = shader.Vertex(face.V3)

			canva.Rasterize(face)
			canva.Outline(face)
		}
	}
}
