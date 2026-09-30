package graph_v1

import (
	"image/color"

	"github.com/gjantsch/graph/cmd/10/models"
)

type Shader struct {
	Distance  float64
	RotationX float64
	RotationY float64
	RotationZ float64
	Color     color.RGBA
}

func NewShader(distance float64, rgba color.RGBA) *Shader {
	return &Shader{
		Distance:  distance,
		RotationX: 0,
		RotationY: 0,
		RotationZ: 0,
		Color:     rgba,
	}
}

func (s *Shader) Vertex(v models.Coordinate) models.Coordinate {
	v = models.RotateX(v, s.RotationX)
	v = models.RotateY(v, s.RotationY)
	v = models.RotateZ(v, s.RotationZ)
	return models.Project(models.TranslateZ(v, s.Distance))
}

func (s *Shader) Fragment(r, g, b float64) color.RGBA {
	return color.RGBA{
		R: uint8(255),
		G: uint8(0),
		B: uint8(0),
		A: uint8(255),
	}
}
