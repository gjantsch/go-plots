package main

import (
	"image/color"
	"math"
	"time"

	"github.com/zserge/fenster"
)

type Dot struct {
	X int
	Y int
	C color.RGBA
	w fenster.Fenster
}

func NewDot(w fenster.Fenster, x, y int, c color.RGBA) Dot {
	return Dot{w: w, X: x, Y: y, C: c}
}

func (d *Dot) setDot(x, y int, c color.RGBA) {

	// Calculate the distance  between
	// the new and the current position
	dx := x - d.X
	dy := y - d.Y

	// Absolute value of dx
	absDx := dx
	if absDx < 0 {
		absDx = -absDx
	}

	// Absolute value of dy
	absDy := dy
	if absDy < 0 {
		absDy = -absDy
	}

	// Determine the number of steps needed for interpolation
	steps := absDx
	// Choose the larger of adx and ady as the number of steps for interpolation
	if absDy > steps {
		steps = absDy
	}
	// Ensure at least one step to avoid division by zero
	if steps == 0 {
		steps = 1
	}

	// Interpolate and set the dots along the path from the current position to the new position
	for i := 0; i <= steps; i++ {
		ix := d.X + dx*i/steps
		iy := d.Y + dy*i/steps
		d.w.Set(ix, iy, c)
	}

	d.w.Set(x, y, c)

	// Update the dot's current position and color
	d.X = x
	d.Y = y
	d.C = c

}

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
		dot := NewDot(w, 0, height/2, color.RGBA{255, 255, 255, 255})
		for x := 0; x < width; x++ {

			// Calculate y coordinate dynamically
			y := int(float64(height)/2 + math.Sin((float64(x)*0.05))*float64(amplitude))

			// Light up a single pixel dot
			dot.setDot(x, y, white)
		}
	}
}
