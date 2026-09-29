package graph_v1

import (
	"image/color"

	"github.com/gjantsch/graph/cmd/09/models"
)

type Shader struct {
	Distance float64
}

func NewShader(distance float64) *Shader {
	return &Shader{
		Distance: distance,
	}
}

func (s *Shader) Vertex(v models.Coordinate) models.Coordinate {
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
