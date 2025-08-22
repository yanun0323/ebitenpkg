package ebitenpkg

// controller controls the movement of an object.
//
// It can be used to move an object, rotate it, scale it, and align it.
//
// controller is thread unsafe.
type controller struct {
	direction        Direction
	align            Align
	animation        Animation
	movement         Vector
	movementAddition chan *controllerDelta[Vector]
	scaled           bool /* for init scale */
	scale            Vector
	scaleAddition    chan *controllerDelta[Vector]
	rotation         float64
	rotationAddition chan *controllerDelta[Vector]
	colored          bool
	color            [4]uint8
	colorAddition    chan *controllerDelta[colors]
}

func (ctr *controller) SetAlign(a Align) {
	ctr.align = a
}

func (ctr *controller) SetAnimation(a Animation) {
	ctr.animation = a
	resetAnimation(ctr.movementAddition, a)
	resetAnimation(ctr.scaleAddition, a)
	resetAnimation(ctr.rotationAddition, a)
	resetAnimation(ctr.colorAddition, a)
}

func (ctr *controller) GetAnimation() Animation {
	return ctr.animation
}

func resetAnimation[T deltaValue[T]](ch chan *controllerDelta[T], a Animation) {
	l := len(ch)
	for i := 0; i < l; i++ {
		c := <-ch
		c.animation = a
		ch <- c
	}
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

	add, rp := newControllerDelta(Vector{x, y}, tick, len(replace) != 0 && replace[0], ctr.animation)
	if rp || ctr.movementAddition == nil {
		ctr.movementAddition = make(chan *controllerDelta[Vector], _defaultChanCap)
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

	add, rp := newControllerDelta(Vector{X: degree}, tick, len(replace) != 0 && replace[0], ctr.animation)
	if rp || ctr.rotationAddition == nil {
		ctr.rotationAddition = make(chan *controllerDelta[Vector], _defaultChanCap)
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

	add, rp := newControllerDelta(Vector{x, y}, tick, len(replace) != 0 && replace[0], ctr.animation)
	if rp || ctr.scaleAddition == nil {
		ctr.scaleAddition = make(chan *controllerDelta[Vector], _defaultChanCap)
	}

	ctr.scaleAddition <- add
}

func (ctr *controller) SetColor(r, g, b, a uint8) {
	if !ctr.colored {
		ctr.color = [4]uint8{255, 255, 255, 255}
		ctr.colored = true
	}

	ctr.color = [4]uint8{r, g, b, a}
}

func (ctr *controller) SetColoring(r, g, b, a uint8, tick int) {
	if tick <= 0 {
		return
	}

	if !ctr.colored {
		ctr.color = [4]uint8{255, 255, 255, 255}
		ctr.colored = true
	}

	add, rp := newControllerDelta(colors{r, g, b, a}, tick, true, ctr.animation)
	if rp || ctr.colorAddition == nil {
		ctr.colorAddition = make(chan *controllerDelta[colors], _defaultChanCap)
	}

	ctr.colorAddition <- add
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

		cache = add.CalculateResult(tick, cache)
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

		cache = add.CalculateResult(tick, Vector{X: cache}).X
		ctr.rotationAddition <- add
	}

	ctr.rotation = cache

	return ctr.rotation
}

func (ctr *controller) GetScale() (x, y float64) {
	if !ctr.scaled {
		return 1, 1
	}
	tick := CurrentGameTime()
	cache := ctr.scale
	for i := len(ctr.scaleAddition) - 1; i >= 0; i-- {
		add := <-ctr.scaleAddition
		if add.IsComplete() {
			continue
		}

		cache = add.CalculateResult(tick, cache)
		ctr.scaleAddition <- add
	}

	ctr.scale = cache
	return ctr.scale.X, ctr.scale.Y
}

func (ctr *controller) GetColor() (r, g, b, a uint8) {
	if !ctr.colored {
		return 255, 255, 255, 255
	}

	tick := CurrentGameTime()
	cache := ctr.color
	for i := len(ctr.colorAddition) - 1; i >= 0; i-- {
		add := <-ctr.colorAddition
		if add.IsComplete() {
			continue
		}

		cache = add.CalculateResult(tick, cache)
		ctr.colorAddition <- add
	}

	ctr.color = cache
	return ctr.color[0], ctr.color[1], ctr.color[2], ctr.color[3]
}

func (ctr *controller) GetBarycenter(parentMovement ...Vector) (float64, float64) {
	if len(parentMovement) == 0 {
		return ctr.movement.X, ctr.movement.Y
	}

	return ctr.movement.X + parentMovement[0].X, ctr.movement.Y + parentMovement[0].Y
}
