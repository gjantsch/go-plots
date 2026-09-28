package main

import (
	"image/color"
	"math"
	"time"

	"github.com/zserge/fenster"
)

func main() {
	width, height := 500, 500
	amplitude := 100

	// 1. Open the tiniest window possible
	w, err := fenster.New(width, height, "Dots Canvas")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// 2. Main polling frame loop (60 FPS target)
	for w.Loop(time.Second / 60) {

		// Clear screen to black every frame
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				w.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
		}

		// 3. Draw graphical forms purely out of individual dots
		// Example: Generating a mathematical sine-wave shape from isolated pixels
		white := color.RGBA{255, 255, 255, 255}
		for x := 0; x < width; x++ {
			// Calculate y coordinate dynamically
			y := int(float64(height)/2 + math.Sin((float64(x)*0.05))*float64(amplitude))

			// Light up a single pixel dot
			w.Set(x, y, white)
		}
	}
}
