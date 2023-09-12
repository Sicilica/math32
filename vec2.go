package math32

// Vec2 represents an offset in 2D space.
type Vec2 struct {
	X, Y float32
}

// Dot returns the dot product of a and b.
func (a Vec2) Dot(b Vec2) float32 {
	return a.X*b.X + a.Y*b.Y
}

// Minus returns the difference between each component of a and b.
func (a Vec2) Minus(b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}

// Normalized returns a vector in the same direction as v with 1.0 length.
//
// Special cases:
//  [0 0] => [0 0]
func (v Vec2) Normalized() Vec2 {
	return v.Times(1 / Sqrt(v.Dot(v)))
}

// Minus returns the sum of each component of a and b.
func (a Vec2) Plus(b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

// ProjectedOnto returns the portion of a that is parallel to b, scaled by the length
// of b.
func (a Vec2) ProjectedOnto(b Vec2) Vec2 {
	return b.Times(a.Dot(b))
}

// Times returns the result of n multiplied by a scalar s.
func (v Vec2) Times(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}
