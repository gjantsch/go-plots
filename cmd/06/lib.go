package main

import (
	"image/color"
	"math"

	"github.com/zserge/fenster"
)

type Vector struct {
	X int
	Y int
}

func MinX(vectors []Vector) int {
	min := vectors[0].X
	for _, v := range vectors {
		if v.X < min {
			min = v.X
		}
	}
	return min
}

func MinY(vectors []Vector) int {
	min := vectors[0].Y
	for _, v := range vectors {
		if v.Y < min {
			min = v.Y
		}
	}
	return min
}

func MaxX(vectors []Vector) int {
	max := vectors[0].X
	for _, v := range vectors {
		if v.X > max {
			max = v.X
		}
	}
	return max
}

func MaxY(vectors []Vector) int {
	max := vectors[0].Y
	for _, v := range vectors {
		if v.Y > max {
			max = v.Y
		}
	}
	return max
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SignedTriangleArea(a, b, c Vector) float64 {
	return 0.5 * float64((b.Y-a.Y)*(b.X+a.X)+(c.Y-b.Y)*(c.X+b.X)+(a.Y-c.Y)*(a.X+c.X))
}

func Line(w fenster.Fenster, origin Vector, destiny Vector, c color.RGBA) {

	x0 := float64(origin.X)
	y0 := float64(origin.Y)
	x1 := float64(destiny.X)
	y1 := float64(destiny.Y)

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

func Triangle(v1, v2, v3 Vector, w fenster.Fenster, c color.RGBA) {
	// the bounding box vertices
	topLeft := Vector{X: MinX([]Vector{v1, v2, v3}), Y: MinY([]Vector{v1, v2, v3})}
	bottomRight := Vector{X: MaxX([]Vector{v1, v2, v3}), Y: MaxY([]Vector{v1, v2, v3})}

	area := SignedTriangleArea(v1, v2, v3)

	// scan the bounding box for dots inside
	for x := topLeft.X; x <= bottomRight.X; x++ {
		for y := topLeft.Y; y <= bottomRight.Y; y++ {
			alfa := SignedTriangleArea(v1, v2, Vector{X: x, Y: y}) / area
			beta := SignedTriangleArea(v2, v3, Vector{X: x, Y: y}) / area
			gamma := SignedTriangleArea(v3, v1, Vector{X: x, Y: y}) / area
			if alfa >= 0 && beta >= 0 && gamma >= 0 {
				w.Set(int(x), int(y), c)
			}
		}
	}

	Line(w, v1, v2, c)
	Line(w, v2, v3, c)
	Line(w, v3, v1, c)

}

func TriangleAlfa(v1, v2, v3 Vector, w fenster.Fenster, c color.RGBA) {
	// the bounding box vertices
	topLeft := Vector{X: MinX([]Vector{v1, v2, v3}), Y: MinY([]Vector{v1, v2, v3})}
	bottomRight := Vector{X: MaxX([]Vector{v1, v2, v3}), Y: MaxY([]Vector{v1, v2, v3})}

	area := SignedTriangleArea(v1, v2, v3)

	// scan the bounding box for dots inside
	for x := topLeft.X; x <= bottomRight.X; x++ {
		for y := topLeft.Y; y <= bottomRight.Y; y++ {
			alfa := SignedTriangleArea(v1, v2, Vector{X: x, Y: y}) / area
			beta := SignedTriangleArea(v2, v3, Vector{X: x, Y: y}) / area
			gamma := SignedTriangleArea(v3, v1, Vector{X: x, Y: y}) / area
			if alfa >= 0 && beta >= 0 && gamma >= 0 {
				r := uint8(math.Round(255 * alfa))
				g := uint8(math.Round(255 * beta))
				b := uint8(math.Round(255 * gamma))
				color := color.RGBA{r, g, b, 255}
				w.Set(int(x), int(y), color)
			}
		}
	}

	// Line(w, v1, v2, c)
	// Line(w, v2, v3, c)
	// Line(w, v3, v1, c)

}
