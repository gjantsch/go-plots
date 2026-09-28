package main

import (
	"image/color"
	"time"

	graph_v1 "github.com/gjantsch/graph/pkg/graph/v1"
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

	greenColor := graph_v1.NewShader(0.5, 1, 0.5, 0.1)
	blueColor := graph_v1.NewShader(1, 1, 1, 1)

	// Open the tiniest window possible
	w, err := fenster.New(Width, Height, "Dots Canvas")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// Build the canvas object using the fenster window
	canva := graph_v1.NewCanva(w, Width, Height)

	// Main polling frame loop (60 FPS target)
	for w.Loop(time.Second / 60) {

		// Clear screen to black every frame
		canva.Clear(color.RGBA{0, 0, 0, 255})

		v1 := graph_v1.Vector{X: 200, Y: 300}
		v2 := graph_v1.Vector{X: 300, Y: 300}
		v3 := graph_v1.Vector{X: 250, Y: 200}
		canva.Rasterize(v1, v2, v3, blueColor)

		v21 := graph_v1.Vector{X: 50, Y: 12}
		v22 := graph_v1.Vector{X: 120, Y: 98}
		v23 := graph_v1.Vector{X: 220, Y: 50}
		canva.Rasterize(v21, v22, v23, greenColor)

	}
}
