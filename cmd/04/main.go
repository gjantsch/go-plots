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

// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
func BresenhamLine(w fenster.Fenster, x0, y0, x1, y1 int, c color.RGBA) {
	dx := x1 - x0
	dy := y1 - y0
	stepX := 1
	stepY := 1

	if dx < 0 {
		dx = -dx
		stepX = -1
	}
	if dy < 0 {
		dy = -dy
		stepY = -1
	}

	err := dx - dy
	for {
		w.Set(x0, y0, c)
		if x0 == x1 && y0 == y1 {
			break
		}

		twiceErr := 2 * err
		if twiceErr > -dy {
			err -= dy
			x0 += stepX
		}
		if twiceErr < dx {
			err += dx
			y0 += stepY
		}
	}

}

func NewDot(w fenster.Fenster, x, y int, c color.RGBA) Dot {
	return Dot{w: w, X: x, Y: y, C: c}
}

func (d *Dot) setDot(x, y int, c color.RGBA) {

	BresenhamLine(d.w, x, y, d.X, d.Y, c)

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
