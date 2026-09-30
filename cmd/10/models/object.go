package models

import (
	"image/color"
	"math"
)

type Coordinate struct {
	X float64
	Y float64
	Z float64
}

func NewCoordinate(x, y, z float64) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
		Z: z,
	}
}

func Project(a Coordinate) Coordinate {
	return Coordinate{
		X: a.X / a.Z,
		Y: a.Y / a.Z,
		Z: a.Z,
	}
}

func TranslateZ(a Coordinate, distance float64) Coordinate {
	return Coordinate{
		X: a.X,
		Y: a.Y,
		Z: a.Z + distance,
	}
}

func RotateX(a Coordinate, angle float64) Coordinate {
	if angle == 0 {
		return a
	}
	return Coordinate{
		X: a.X,
		Y: a.Y*math.Cos(angle) - a.Z*math.Sin(angle),
		Z: a.Y*math.Sin(angle) + a.Z*math.Cos(angle),
	}
}

func RotateZ(a Coordinate, angle float64) Coordinate {
	if angle == 0 {
		return a
	}
	// Implement rotation logic here (e.g., around the Z-axis)
	return Coordinate{
		X: a.X*math.Cos(angle) - a.Y*math.Sin(angle),
		Y: a.X*math.Sin(angle) + a.Y*math.Cos(angle),
		Z: a.Z,
	}
}

func RotateY(a Coordinate, angle float64) Coordinate {
	if angle == 0 {
		return a
	}
	return Coordinate{
		X: a.X*math.Cos(angle) + a.Z*math.Sin(angle),
		Y: a.Y,
		Z: -a.X*math.Sin(angle) + a.Z*math.Cos(angle),
	}
}

type Face struct {
	V1  Coordinate
	V2  Coordinate
	V3  Coordinate
	RGB color.RGBA
}

func NewFace(v1, v2, v3 Coordinate, rgb color.RGBA) Face {
	return Face{
		V1:  v1,
		V2:  v2,
		V3:  v3,
		RGB: rgb,
	}
}

type Object struct {
	Faces    []Face
	depth    float64
	rotation float64
}

func NewObject(faces []Face, depth float64, rotation float64) Object {
	return Object{
		Faces:    faces,
		depth:    depth,
		rotation: rotation,
	}
}

func (o *Object) AddFace(f Face) {
	o.Faces = append(o.Faces, f)
}
