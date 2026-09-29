package models

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

func (a *Coordinate) Project() Coordinate {
	return Coordinate{
		X: a.X / a.Z,
		Y: a.Y / a.Z,
		Z: a.Z,
	}
}

func (a *Coordinate) TranslateZ(distance float64) Coordinate {
	return Coordinate{
		X: a.X,
		Y: a.Y,
		Z: a.Z + distance,
	}
}

type Face struct {
	V1 Coordinate
	V2 Coordinate
	V3 Coordinate
}

func NewFace(v1, v2, v3 Coordinate) Face {
	return Face{
		V1: v1,
		V2: v2,
		V3: v3,
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
