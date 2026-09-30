package graph_v1

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

func SignedTriangleArea(a, b, c Vector) float64 {
	return 0.5 * float64((b.Y-a.Y)*(b.X+a.X)+(c.Y-b.Y)*(c.X+b.X)+(a.Y-c.Y)*(a.X+c.X))
}
