package math32

// Rect represents an axis-aligned box in 2d space.
type Rect struct {
	X0, Y0, X1, Y1 float32
}

// Overlaps returns true if a and b overlap.
//
// Special cases:
//  If a and b share an edge, they are counted as overlapping.
func (a Rect) Overlaps(b Rect) bool {
	return !(a.X0 > b.X1 || a.X1 < b.X0 || a.Y0 > b.Y1 || a.Y1 < b.Y0)
}

// Translated returns the result of adding a translation to every point of the rect.
func (r Rect) Translated(t Vec2) Rect {
	return Rect{r.X0 + t.X, r.Y0 + t.Y, r.X1 + t.X, r.Y1 + t.Y}
}
