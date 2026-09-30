package graph_v1

import (
	"image/color"
	"math"

	"github.com/gjantsch/graph/cmd/09/models"
	"github.com/zserge/fenster"
)

type Canvas struct {
	Width  int
	Height int
	canva  fenster.Fenster
}

func NewCanva(f fenster.Fenster, width, height int) *Canvas {
	return &Canvas{
		Width:  width,
		Height: height,
		canva:  f,
	}
}

// Draw a line from origin to destiny with the specified color
func (c *Canvas) Line(origin Vector, destiny Vector, col color.RGBA) {

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

	c.Set(math.Round(x0), math.Round(y0), col)

	for x0 != x1 || y0 != y1 {
		x0 += stepX
		y0 += stepY

		if (stepX > 0 && x0 > x1) || (stepX < 0 && x0 < x1) {
			x0 = x1
		}

		if (stepY > 0 && y0 > y1) || (stepY < 0 && y0 < y1) {
			y0 = y1
		}
		c.Set(math.Round(x0), math.Round(y0), col)
	}
}

func (c *Canvas) Rasterize(face models.Face) {
	v1, v2, v3 := c.FaceToScreen(face)

	rgb := color.RGBA{R: 255, G: 0, B: 0, A: 255}

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
				c.canva.Set(x, y, rgb)
			}
		}
	}

	c.Line(v1, v2, rgb)
	c.Line(v2, v3, rgb)
	c.Line(v3, v1, rgb)
}

func (c *Canvas) Set(x, y float64, col color.RGBA) {
	c.canva.Set(int(math.Round(x)), int(math.Round(y)), col)
}

func (c *Canvas) Clear(col color.RGBA) {
	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			c.canva.Set(x, y, col)
		}
	}
}

func (c *Canvas) FaceToScreen(f models.Face) (Vector, Vector, Vector) {
	return c.Screen(f.V1), c.Screen(f.V2), c.Screen(f.V3)
}

func (c *Canvas) Screen(f models.Coordinate) Vector {
	return Vector{
		X: int((f.X + 1) * float64(c.Width) / 2),
		Y: int((f.Y + 1) * float64(c.Height) / 2),
	}
}
