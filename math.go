package ebitenpkg

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func isOverlap(a, b []Vector) bool {
	for _, v := range a {
		for _, w := range b {
			if isInside(a, w) || isInside(b, v) {
				return true
			}
		}
	}

	return false
}

func isInside(area []Vector, p Vector) bool {
	switch len(area) {
	case 3:
		return pointInTriangle(area[0], area[1], area[2], p)
	case 4:
		return pointInPolygon(area[0], area[1], area[2], area[3], p)
	default:
		return false
	}
}

func pointInPolygon(a, b, c, d, p Vector) bool {
	area := polygonArea(a, b, c, d) + _floatFix

	pAreaSummed := triangleArea(a, b, p) +
		triangleArea(b, c, p) +
		triangleArea(c, d, p) +
		triangleArea(d, a, p)

	return pAreaSummed <= area
}

func pointInTriangle(a, b, c, p Vector) bool {
	area := triangleArea(a, b, c) + _floatFix

	pAreaSummed := triangleArea(a, b, p) +
		triangleArea(b, c, p) +
		triangleArea(a, c, p)

	return pAreaSummed <= area
}

func polygonArea(p1, p2, p3, p4 Vector) float64 {
	return 0.5 * math.Abs(
		p1.X*p2.Y+p2.X*p3.Y+p3.X*p4.Y+p4.X*p1.Y-
			p1.Y*p2.X-p2.Y*p3.X-p3.Y*p4.X-p4.Y*p1.X,
	)
}

func triangleArea(a, b, c Vector) float64 {
	return 0.5 * math.Abs(a.X*(b.Y-c.Y)+b.X*(c.Y-a.Y)+c.X*(a.Y-b.Y))
}

func rotateVector(center, target Vector, degree float64) Vector {
	// Convert the angle from degrees to radians
	theta := degree * math.Pi / 180
	dX := target.X - center.X
	dY := target.Y - center.Y
	// Calculate the new coordinates of point Y
	target.X = dX*math.Cos(theta) - dY*math.Sin(theta) + center.X
	target.Y = dX*math.Sin(theta) + dY*math.Cos(theta) + center.Y

	return target
}

func scaleVector(center, target, scale Vector) Vector {
	switch scale.X {
	case 0:
		target.X = 0
	case 1:
	case -1:
		target.X = -target.X
	default:
		target.X = center.X + ((target.X - center.X) * scale.X)
	}

	switch scale.Y {
	case 0:
		target.Y = 0
	case 1:
	case -1:
		target.Y = -target.Y
	default:
		target.Y = center.Y + ((target.Y - center.Y) * scale.Y)
	}

	return target
}

func getVertexes(w, h float64, ctr coords, pr ...attachable) []Vector {
	result := ctr.Aligned().vertexRatio()

	mX, mY := ctr.Moved()
	sX, sY := ctr.Scaled()

	var pmX, pmY float64

	hasParent := len(pr) != 0 && pr[0] != nil

	if hasParent {
		pmX, pmY = pr[0].Moved()

		psX, psY := pr[0].Scaled()
		if psX < 0 {
			mX = -mX
		}

		if psY < 0 {
			mY = -mY
		}
	}

	for i, v := range result {
		v.X *= w
		v.Y *= h

		v = scaleVector(Vector{}, v, Vector{X: sX, Y: sY})
		v = rotateVector(Vector{}, v, ctr.Rotated())

		if hasParent {
			v.X += pmX
			v.Y += pmY
		}

		v.X += mX
		v.Y += mY

		result[i] = v
	}

	return result
}

func getDrawOption(w, h int, current controller, tempScaleX, tempScaleY float64, pr ...attachable) *ebiten.DrawImageOptions {
	mX, mY := current.GetMove()
	oX, oY := current.GetAlign().barycenterOffset(float64(w), float64(h))
	sX, sY := current.GetScale()

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-oX, -oY)
	opt.GeoM.Scale(sX, sY)
	opt.GeoM.Scale(tempScaleX, tempScaleY)
	opt.GeoM.Rotate(current.GetRotate() / _radianToDegree)

	if len(pr) != 0 && pr[0] != nil {
		pmX, pmY := pr[0].Moved()
		opt.GeoM.Translate(pmX, pmY)

		psX, psY := pr[0].Scaled()
		if psX < 0 {
			mX = -mX
		}

		if psY < 0 {
			mY = -mY
		}
	}

	opt.GeoM.Translate(mX, mY)
	return opt
}

/*
	GJK Algorithm
*/

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
