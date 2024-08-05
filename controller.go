package ebitenpkg

// controller controls the movement of an object.
//
// It can be used to move an object, rotate it, scale it, and align it.
//
// controller is thread unsafe.
type controller struct {
	direction Direction
	align     Align
	movement  Vector
	scaled    bool /* for init scale */
	scale     Vector
	rotation  float64
}

func (ctr *controller) SetAlign(a Align) {
	ctr.align = a
}

func (ctr *controller) SetMove(x, y float64, replace ...bool) {
	if len(replace) != 0 && replace[0] {
		ctr.direction = newDirectionFrom(ctr.movement.X, ctr.movement.Y, x, y)
		ctr.movement = Vector{X: x, Y: y}
	} else {
		ctr.direction = newDirection(x, y)
		ctr.movement = Vector{X: ctr.movement.X + x, Y: ctr.movement.Y + y}
	}
}

func (ctr *controller) SetRotate(degree float64, replace ...bool) {
	if len(replace) != 0 && replace[0] {
		ctr.rotation = degree
	} else {
		ctr.rotation += degree
	}
}

func (ctr *controller) SetScale(x, y float64, replace ...bool) {
	if !ctr.scaled {
		ctr.scale = Vector{X: 1, Y: 1}
		ctr.scaled = true
	}

	if len(replace) != 0 && replace[0] {
		ctr.scale = Vector{X: x, Y: y}
	} else {
		ctr.scale = Vector{X: ctr.scale.X * x, Y: ctr.scale.Y * y}
	}
}

func (ctr controller) GetAlign() Align {
	return ctr.align
}

func (ctr controller) GetMove() (x, y float64) {
	return ctr.movement.X, ctr.movement.Y
}

func (ctr controller) GetDirection() Direction {
	return ctr.direction
}

func (ctr controller) GetRotate() float64 {
	return ctr.rotation
}

func (ctr controller) GetScale() (x, y float64) {
	if ctr.scaled {
		return ctr.scale.X, ctr.scale.Y
	}

	return 1, 1
}

func (ctr controller) GetBarycenter(parentMovement ...Vector) (float64, float64) {
	if len(parentMovement) == 0 {
		return ctr.movement.X, ctr.movement.Y
	}

	return ctr.movement.X + parentMovement[0].X, ctr.movement.Y + parentMovement[0].Y
}
