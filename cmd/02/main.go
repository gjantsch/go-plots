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

// this is a forward only version where
// the NewDot happens when x cycled (x==0)
func (d *Dot) setDot(x, y int, c color.RGBA) {

	var startX, endX, startY, endY int

	if x <= d.X {
		startX = x
		endX = d.X
	} else {
		startX = d.X
		endX = x
	}

	if y <= d.Y {
		startY = y
		endY = d.Y
	} else {
		startY = d.Y
		endY = y
	}

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			d.w.Set(i, j, c)
		}
	}

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
		var dot Dot
		for x := 0; x < width; x++ {

			if x == 0 {
				dot = NewDot(w, 0, height/2, white)
			}
			// Calculate y coordinate dynamically
			y := int(float64(height)/2 + math.Sin((float64(x)*0.05))*float64(amplitude))

			// Light up a single pixel dot
			dot.setDot(x, y, white)
		}
	}
}
