package main

import (
	"image/color"
	"time"

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

	greenColor := color.RGBA{0, 255, 0, 255}
	blueColor := color.RGBA{255, 0, 0, 255}

	// 1. Open the tiniest window possible
	w, err := fenster.New(Width, Height, "Dots Canvas")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// 2. Main polling frame loop (60 FPS target)
	for w.Loop(time.Second / 60) {

		// Clear screen to black every frame
		for x := 0; x < Width; x++ {
			for y := 0; y < Height; y++ {
				w.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
		}

		v1 := Vector{X: 200, Y: 300}
		v2 := Vector{X: 300, Y: 300}
		v3 := Vector{X: 250, Y: 200}
		Triangle(v1, v2, v3, w, blueColor)

		v21 := Vector{X: 50, Y: 12}
		v22 := Vector{X: 120, Y: 98}
		v23 := Vector{X: 220, Y: 50}
		TriangleAlfa(v21, v22, v23, w, greenColor)

	}
}
