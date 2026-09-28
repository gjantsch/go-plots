package main

import (
	"image/color"
	"math"
	"time"

	"github.com/zserge/fenster"
)

const (
	Width     = 500
	Height    = 250
	XMin      = 0
	YMin      = 0
	XMax      = Width - 1
	YMax      = Height - 1
	Amplitude = 100
)

func main() {

	xAxisColor := color.RGBA{0, 0, 255, 255}
	yAxisColor := color.RGBA{0, 0, 255, 255}
	diagonalColor := color.RGBA{0, 255, 0, 255}
	whiteColor := color.RGBA{255, 255, 255, 255}
	dimColor := color.RGBA{218, 242, 243, 204}

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

		// draw some pure lines
		midX := int(Width / 2)
		midY := int(Height / 2)

		// Draw axis and diagonals
		// these are for reference and for testing the line drawing algorithm
		Line(w, midX, YMin, midX, YMax, xAxisColor)
		Line(w, XMin, midY, XMax, midY, yAxisColor)
		Line(w, XMin, YMin, XMax, YMax, diagonalColor)
		Line(w, XMin, YMax, XMax, YMin, diagonalColor)

		for x := 0; x < Width; x += XMax / 10 {
			Line(w, x, YMin, x, YMax, dimColor)
		}
		for y := 0; y < Height; y += YMax / 10 {
			Line(w, XMin, y, XMax, y, dimColor)
		}

		// 3. Draw graphical forms purely out of individual dots
		// Example: Generating a mathematical sine-wave shape from isolated pixels
		dot := NewDot(w, 0, midY, whiteColor)
		for x := 0; x < Width; x++ {

			// Calculate y coordinate dynamically
			y := int(float64(Height)/2 + math.Sin((float64(x)*0.05))*float64(Amplitude))

			// Light up a single pixel dot
			dot.setDot(x, y, whiteColor)
		}
	}
}
