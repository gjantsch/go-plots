package main

import (
	"image/color"
	"math"

	"github.com/zserge/fenster"
)

type Dot struct {
	X int
	Y int
	C color.RGBA
	w fenster.Fenster
}

func Line(w fenster.Fenster, xO, yO, xD, yD int, c color.RGBA) {

	x0 := float64(xO)
	y0 := float64(yO)
	x1 := float64(xD)
	y1 := float64(yD)

	dx := x1 - x0
	dy := y1 - y0
	stepX := 1.0
	stepY := 1.0

	// Adjust absolute values and direction
	if dx < 0 {
		dx = -dx
		stepX = -1
	}

	if dy < 0 {
		dy = -dy
		stepY = -1
	}

	// who is the biggest delta
	if dx > dy {
		if dx != 0 {
			stepY = stepY * dy / dx
		}
	} else {
		if dy != 0 {
			stepX = stepX * dx / dy
		}
	}

	w.Set(int(math.Round(x0)), int(math.Round(y0)), c)

	for x0 != x1 || y0 != y1 {
		x0 += stepX
		y0 += stepY

		if (stepX > 0 && x0 > x1) || (stepX < 0 && x0 < x1) {
			x0 = x1
		}

		if (stepY > 0 && y0 > y1) || (stepY < 0 && y0 < y1) {
			y0 = y1
		}
		w.Set(int(math.Round(x0)), int(math.Round(y0)), c)
	}
}

func NewDot(w fenster.Fenster, x, y int, c color.RGBA) Dot {
	return Dot{w: w, X: x, Y: y, C: c}
}

func (d *Dot) setDot(x, y int, c color.RGBA) {

	Line(d.w, d.X, d.Y, x, y, c)

	// Update the dot's current position and color
	d.X = x
	d.Y = y
	d.C = c

}
