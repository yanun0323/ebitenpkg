package ebitenpkg

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var _floatFix float64 = 0.001

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

type parent struct {
	w, h float64
	ctr  Controllable[any]
}

func getVertexes(w, h float64, current controller, pr ...parent) []Vector {
	result := current.Aligned().vertexRatio()

	mX, mY := current.Moved()
	sX, sY := current.Scaled()

	var (
		pmX, pmY float64
		poX, poY float64
		psX, psY float64
		pR       float64
	)

	hasParent := len(pr) != 0 && pr[0].ctr != nil

	if hasParent {
		pmX, pmY = pr[0].ctr.Moved()
		poX, poY = pr[0].ctr.Aligned().barycenterOffset(pr[0].w, pr[0].h)
		psX, psY = pr[0].ctr.Scaled()
		pR = pr[0].ctr.Rotated()
	}

	for i, v := range result {
		v.X *= w
		v.Y *= h

		v = scaleVector(Vector{}, v, Vector{X: sX, Y: sY})
		v = rotateVector(Vector{}, v, current.Rotated())

		if hasParent {
			v.X = v.X + pmX - poX
			v.Y = v.Y + pmY - poY
			v = scaleVector(Vector{}, v, Vector{X: psX, Y: psY})
			v = rotateVector(Vector{}, v, pR)
		}

		v.X += mX
		v.Y += mY

		result[i] = v
	}

	return result
}

func getDrawOption(w, h float64, current controller, pr ...parent) *ebiten.DrawImageOptions {
	mX, mY := current.Moved()
	oX, oY := current.Aligned().barycenterOffset(w, h)
	sX, sY := current.Scaled()

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-oX, -oY)
	opt.GeoM.Scale(sX, sY)
	opt.GeoM.Rotate(current.Rotated() / radianToDegree)

	if len(pr) != 0 && pr[0].ctr != nil {
		pmX, pmY := pr[0].ctr.Moved()
		poX, poY := pr[0].ctr.Aligned().barycenterOffset(pr[0].w, pr[0].h)
		psX, psY := pr[0].ctr.Scaled()
		pR := pr[0].ctr.Rotated()
		opt.GeoM.Translate(pmX-poX, pmY-poY)
		opt.GeoM.Scale(psX, psY)
		opt.GeoM.Rotate(pR / radianToDegree)
	}

	opt.GeoM.Translate(mX, mY)
	return opt
}
