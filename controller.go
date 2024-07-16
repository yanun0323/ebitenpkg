package ebitenpkg

import (
	"math"
)

var radianToDegree = 180 / math.Pi

type controller struct {
	movement Vector
	scale    Vector
	rotation float64
	align    Align
}

func newController(a ...Align) controller {
	align := AlignTopLeading
	if len(a) != 0 {
		align = a[0]
	}

	return controller{
		movement: Vector{X: 0, Y: 0},
		scale:    Vector{X: 1, Y: 1},
		align:    align,
	}
}

func (ctr *controller) Align(a Align) *controller {
	ctr.align = a
	return ctr
}

func (ctr *controller) Move(x, y float64, replace ...bool) *controller {
	if len(replace) != 0 && replace[0] {
		ctr.movement = Vector{X: x, Y: y}
	} else {
		ctr.movement = Vector{X: ctr.movement.X + x, Y: ctr.movement.Y + y}
	}

	return ctr
}

func (ctr *controller) Rotate(degree float64, replace ...bool) *controller {
	if len(replace) != 0 && replace[0] {
		ctr.rotation = degree
	} else {
		ctr.rotation += degree
	}

	return ctr
}

func (ctr *controller) Scale(x, y float64, replace ...bool) *controller {
	if len(replace) != 0 && replace[0] {
		ctr.scale = Vector{X: x, Y: y}
	} else {
		ctr.scale = Vector{X: ctr.scale.X * x, Y: ctr.scale.Y * y}
	}

	return ctr
}

func (ctr controller) Aligned() Align {
	return ctr.align
}

func (ctr controller) Moved() (x, y float64) {
	return ctr.movement.X, ctr.movement.Y
}

func (ctr controller) Rotated() float64 {
	return ctr.rotation
}

func (ctr controller) Scaled() (x, y float64) {
	return ctr.scale.X, ctr.scale.Y
}

func (ctr controller) Copy() controller {
	return ctr
}

func (ctr controller) Barycenter(parent ...Controllable[any]) (float64, float64) {
	if len(parent) == 0 || parent[0] == nil {
		return ctr.movement.X, ctr.movement.Y
	}

	pX, pY := parent[0].Moved()
	return ctr.movement.X + pX, ctr.movement.Y + pY
}
