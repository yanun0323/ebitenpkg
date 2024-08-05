package ebitenpkg

import (
	"math"
)

// dot product of two vectors
func dot(a, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y
}

// Subtract two vectors
func sub(a, b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}

// Negate a vector
func neg(a Vector) Vector {
	return Vector{-a.X, -a.Y}
}

// Triple product
func tripleProduct(a, b, c Vector) Vector {
	ac := dot(a, c)
	bc := dot(b, c)
	return Vector{b.X*ac - a.X*bc, b.Y*ac - a.Y*bc}
}

// Find the farthest point in a given direction in a shape
func support(shape []Vector, dir Vector) Vector {
	maxDot := -math.MaxFloat64
	farthest := shape[0]
	for _, v := range shape {
		d := dot(v, dir)
		if d > maxDot {
			maxDot = d
			farthest = v
		}
	}
	return farthest
}

// Check if the origin is in the simplex and update the direction
func nextSimplex(simplex []Vector, dir *Vector) bool {
	a := simplex[len(simplex)-1]
	ao := neg(a)
	if len(simplex) == 3 {
		b := simplex[1]
		c := simplex[0]
		ab := sub(b, a)
		ac := sub(c, a)
		abPerp := tripleProduct(ac, ab, ab)
		acPerp := tripleProduct(ab, ac, ac)
		if dot(abPerp, ao) > -_floatFix {
			*dir = abPerp
		} else {
			if dot(acPerp, ao) > _floatFix {
				*dir = acPerp
			} else {
				return true
			}
		}
	} else {
		b := simplex[0]
		ab := sub(b, a)
		*dir = tripleProduct(ab, ao, ab)
		if dir.X <= _floatFix && dir.Y <= _floatFix {
			*dir = Vector{-ab.Y, ab.X}
		}
	}
	return false
}

// gjk algorithm to check if two shapes intersect
func gjk(shapeA, shapeB []Vector) bool {
	dir := Vector{1, 1}
	simplex := []Vector{sub(support(shapeA, dir), support(shapeB, neg(dir)))}
	dir = neg(simplex[0])

	for i := 0; i < 5; i++ {
		a := sub(support(shapeA, dir), support(shapeB, neg(dir)))
		if dot(a, dir) <= _floatFix {
			return false
		}

		simplex = append(simplex, a)
		if nextSimplex(simplex, &dir) {
			return true
		}
	}

	return false
}
