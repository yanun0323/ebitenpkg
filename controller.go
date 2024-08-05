package ebitenpkg

type controllerAdditionValue struct {
	targetTick int
	lastTick   int
	offset     Vector
	replace    bool
}

// controller controls the movement of an object.
//
// It can be used to move an object, rotate it, scale it, and align it.
//
// controller is thread unsafe.
type controller struct {
	direction        Direction
	align            Align
	movement         Vector
	movementAddition map[controllerAdditionValue]struct{}
	scaled           bool /* for init scale */
	scale            Vector
	scaleAddition    map[controllerAdditionValue]struct{}
	rotation         float64
	rotationAddition map[controllerAdditionValue]struct{}
}

func (ctr *controller) SetAlign(a Align) {
	ctr.align = a
}

func (ctr *controller) SetMove(x, y float64, replace ...bool) {
	for add := range ctr.movementAddition {
		if add.replace && add.targetTick < add.lastTick {
			return
		}
	}

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

	offset := Vector{X: x / float64(tick), Y: y / float64(tick)}
	rp := len(replace) != 0 && replace[0]
	if rp {
		offset = Vector{X: ctr.movement.X + offset.X, Y: ctr.movement.Y + offset.Y}
	}

	current := CurrentGameTime()
	add := controllerAdditionValue{
		targetTick: current + tick,
		lastTick:   current,
		offset:     offset,
		replace:    rp,
	}

	if rp || ctr.movementAddition == nil {
		ctr.movementAddition = map[controllerAdditionValue]struct{}{add: {}}
	} else {
		ctr.movementAddition[add] = struct{}{}
	}
}

func (ctr *controller) SetRotate(degree float64, replace ...bool) {
	for add := range ctr.rotationAddition {
		if add.replace && add.targetTick < add.lastTick {
			return
		}
	}

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

	rp := len(replace) != 0 && replace[0]
	current := CurrentGameTime()
	add := controllerAdditionValue{
		targetTick: current + tick,
		lastTick:   current,
		offset:     Vector{X: degree / float64(tick), Y: 0},
		replace:    rp,
	}

	if rp || ctr.rotationAddition == nil {
		ctr.rotationAddition = map[controllerAdditionValue]struct{}{add: {}}
	} else {
		ctr.rotationAddition[add] = struct{}{}
	}
}

func (ctr *controller) SetScale(x, y float64, replace ...bool) {
	for add := range ctr.scaleAddition {
		if add.replace && add.targetTick < add.lastTick {
			return
		}
	}

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

	offset := Vector{X: x / float64(tick), Y: y / float64(tick)}
	rp := len(replace) != 0 && replace[0]
	if rp {
		offset = Vector{X: ctr.scale.X * offset.X, Y: ctr.scale.Y * offset.Y}
	}

	current := CurrentGameTime()
	add := controllerAdditionValue{
		targetTick: current + tick,
		lastTick:   current,
		offset:     offset,
		replace:    rp,
	}

	if rp || ctr.scaleAddition == nil {
		ctr.scaleAddition = map[controllerAdditionValue]struct{}{add: {}}
	} else {
		ctr.scaleAddition[add] = struct{}{}
	}
}

func (ctr controller) GetAlign() Align {
	return ctr.align
}

func (ctr *controller) GetMove() (x, y float64) {
	tick := CurrentGameTime()
	for add := range ctr.movementAddition {
		if add.targetTick <= add.lastTick {
			delete(ctr.movementAddition, add)
			continue
		}

		var tickOffset float64
		if tick >= add.targetTick {
			tickOffset = float64(add.targetTick - add.lastTick)
		} else {
			tickOffset = float64(tick - add.lastTick)
		}

		if add.replace {
			ctr.movement = Vector{X: add.offset.X, Y: add.offset.Y}
		} else {
			valueX := add.offset.X * tickOffset
			valueY := add.offset.Y * tickOffset
			ctr.movement = Vector{X: ctr.movement.X + valueX, Y: ctr.movement.Y + valueY}
		}

		add.lastTick = tick
	}

	return ctr.movement.X, ctr.movement.Y
}

func (ctr controller) GetDirection() Direction {
	return ctr.direction
}

func (ctr *controller) GetRotate() float64 {
	tick := CurrentGameTime()
	for add := range ctr.rotationAddition {
		if add.targetTick <= add.lastTick {
			delete(ctr.rotationAddition, add)
			continue
		}

		var tickOffset float64
		if tick >= add.targetTick {
			tickOffset = float64(add.targetTick - add.lastTick)
		} else {
			tickOffset = float64(tick - add.lastTick)
		}

		value := add.offset.X * tickOffset

		if add.replace {
			ctr.rotation = value
		} else {
			ctr.rotation += value
		}

		add.lastTick = tick
	}

	return ctr.rotation
}

func (ctr *controller) GetScale() (x, y float64) {
	tick := CurrentGameTime()
	for add := range ctr.scaleAddition {
		if add.targetTick <= add.lastTick {
			delete(ctr.scaleAddition, add)
			continue
		}

		var tickOffset float64
		if tick >= add.targetTick {
			tickOffset = float64(add.targetTick - add.lastTick)
		} else {
			tickOffset = float64(tick - add.lastTick)
		}

		if add.replace {
			ctr.scale = Vector{X: add.offset.X, Y: add.offset.Y}
		} else {
			valueX := add.offset.X * tickOffset
			valueY := add.offset.Y * tickOffset
			ctr.scale = Vector{X: ctr.scale.X + valueX, Y: ctr.scale.Y + valueY}
		}

		add.lastTick = tick
	}

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
