package ebitenpkg

// controller controls the movement of an object.
//
// It can be used to move an object, rotate it, scale it, and align it.
//
// controller is thread unsafe.
type controller struct {
	direction        Direction
	align            Align
	movement         Vector
	movementAddition chan controllerDelta
	scaled           bool /* for init scale */
	scale            Vector
	scaleAddition    chan controllerDelta
	rotation         float64
	rotationAddition chan controllerDelta
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

func (ctr *controller) SetMoving(x, y float64, tick int, replace ...bool) {
	if tick <= 0 {
		return
	}

	add, rp := newControllerDelta(x, y, tick, len(replace) != 0 && replace[0])
	if rp || ctr.movementAddition == nil {
		ctr.movementAddition = make(chan controllerDelta, _defaultChanCap)
	}

	ctr.movementAddition <- add
}

func (ctr *controller) SetRotate(degree float64, replace ...bool) {
	if len(replace) != 0 && replace[0] {
		ctr.rotation = degree
	} else {
		ctr.rotation += degree
	}
}

func (ctr *controller) SetRotating(degree float64, tick int, replace ...bool) {
	if tick <= 0 {
		return
	}

	add, rp := newControllerDelta(degree, 0, tick, len(replace) != 0 && replace[0])
	if rp || ctr.rotationAddition == nil {
		ctr.rotationAddition = make(chan controllerDelta, _defaultChanCap)
	}

	ctr.rotationAddition <- add
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

func (ctr *controller) SetScaling(x, y float64, tick int, replace ...bool) {
	if tick <= 0 {
		return
	}

	if !ctr.scaled {
		ctr.scale = Vector{X: 1, Y: 1}
		ctr.scaled = true
	}

	add, rp := newControllerDelta(x, y, tick, len(replace) != 0 && replace[0])
	if rp || ctr.scaleAddition == nil {
		ctr.scaleAddition = make(chan controllerDelta, _defaultChanCap)
	}

	ctr.scaleAddition <- add
}

func (ctr *controller) GetAlign() Align {
	return ctr.align
}

func (ctr *controller) GetMove() (x, y float64) {
	tick := CurrentGameTime()
	cache := ctr.movement
	for i := len(ctr.movementAddition) - 1; i >= 0; i-- {
		add := <-ctr.movementAddition
		if add.IsComplete() {
			continue
		}

		cache = add.CalculateResult(tick, cache, false)
		ctr.movementAddition <- add
	}

	ctr.movement.X, ctr.movement.Y = cache.X, cache.Y

	return ctr.movement.X, ctr.movement.Y
}

func (ctr *controller) GetDirection() Direction {
	return ctr.direction
}

func (ctr *controller) GetRotate() float64 {
	tick := CurrentGameTime()
	cache := ctr.rotation
	for i := len(ctr.rotationAddition) - 1; i >= 0; i-- {
		add := <-ctr.rotationAddition
		if add.IsComplete() {
			continue
		}

		cache = add.CalculateResult(tick, Vector{X: cache}, false).X
		ctr.rotationAddition <- add
	}

	ctr.rotation = cache

	return ctr.rotation
}

func (ctr *controller) GetScale() (x, y float64) {
	tick := CurrentGameTime()
	cache := ctr.scale
	for i := len(ctr.scaleAddition) - 1; i >= 0; i-- {
		add := <-ctr.scaleAddition
		if add.IsComplete() {
			continue
		}

		cache = add.CalculateResult(tick, cache, true)
		ctr.scaleAddition <- add
	}

	ctr.scale = cache
	if ctr.scaled {
		return ctr.scale.X, ctr.scale.Y
	}

	return 1, 1
}

func (ctr *controller) GetBarycenter(parentMovement ...Vector) (float64, float64) {
	if len(parentMovement) == 0 {
		return ctr.movement.X, ctr.movement.Y
	}

	return ctr.movement.X + parentMovement[0].X, ctr.movement.Y + parentMovement[0].Y
}
