package triangle

import (
	"math"
	"sort"
)

// Kind of triangle
type Kind int

const (
	// NaT defines Not a Triangle.
	NaT = iota
	// Equ defines Equilateral Triangle.
	Equ
	// Iso defines Isosceles Triangle.
	Iso
	// Sca define Scalene Triangle.
	Sca
)

type triangle struct {
	a float64
	b float64
	c float64
}

func (t *triangle) isNaN() bool {
	return math.IsNaN(t.a) || math.IsNaN(t.b) || math.IsNaN(t.c)
}

func (t *triangle) isInf() bool {
	return math.IsInf(t.a, 0) || math.IsInf(t.b, 0) || math.IsInf(t.c, 0)
}

func (t *triangle) hasNegative() bool {
	return t.a <= 0 || t.b <= 0 || t.c <= 0
}

func (t *triangle) isDegenerate() bool {
	sides := []float64{t.a, t.b, t.c}
	sort.Float64s(sides)
	return sides[0]+sides[1] < sides[2]
}

func (t *triangle) isNaT() bool {
	return t.isInf() || t.isNaN() || t.isDegenerate() || t.hasNegative()
}

func (t *triangle) isEqu() bool {
	return t.a == t.b && t.b == t.c
}

func (t *triangle) isSca() bool {
	return t.a != t.b && t.a != t.c && t.b != t.c
}

func (t *triangle) isIso() bool {
	return t.a == t.b || t.b == t.c || t.c == t.a
}

// KindFromSides determines the type of Triangle.
func KindFromSides(a, b, c float64) Kind {
	triangle := &triangle{a, b, c}
	if triangle.isNaT() {
		return NaT
	}
	if triangle.isEqu() {
		return Equ
	}
	if triangle.isSca() {
		return Sca
	}
	if triangle.isIso() {
		return Iso
	}
	return NaT
}
