package graph_v1

import "image/color"

type Shader struct {
	r, g, b, alfa float64
}

func NewShader(r, g, b, alfa float64) *Shader {
	return &Shader{
		r:    r,
		g:    g,
		b:    b,
		alfa: alfa,
	}
}

func (s *Shader) RGBA() color.RGBA {
	return color.RGBA{
		R: uint8(s.r * 255),
		G: uint8(s.g * 255),
		B: uint8(s.b * 255),
		A: uint8(s.alfa * 255),
	}
}

func (s *Shader) Fragment(r, g, b float64) color.RGBA {
	return color.RGBA{
		R: uint8(s.r * r * 255),
		G: uint8(s.g * g * 255),
		B: uint8(s.b * b * 255),
		A: uint8(s.alfa * 255),
	}
}
