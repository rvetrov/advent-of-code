package geom

type PointInt struct {
	X, Y, Z int
}

func (p PointInt) Add(v VectorInt) PointInt {
	return PointInt{X: p.X + v.X, Y: p.Y + v.Y, Z: p.Z + v.Z}
}

func (p PointInt) Sub(rhs PointInt) VectorInt {
	return VectorInt{X: p.X - rhs.X, Y: p.Y - rhs.Y, Z: p.Z - rhs.Z}
}

type PointFloat struct {
	X, Y, Z float64
}

type VectorInt struct {
	X, Y, Z int
}

func (v VectorInt) Multiply(m int) VectorInt {
	return VectorInt{X: v.X * m, Y: v.Y * m, Z: v.Z * m}
}

func DotProduct(v1, v2 VectorInt) int {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}
