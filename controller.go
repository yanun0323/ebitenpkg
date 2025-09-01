package ebitenpkg

// controller controls the movement of an object.
//
// It can be used to move an object, rotate it, scale it, and align it.
//
// controller is thread unsafe.
type controller struct {
	direction       Direction
	align           Align
	animation       Animation
	movement        Vector
	movementManager *animationManager[Vector]
	scaled          bool /* for init scale */
	scale           Vector
	scaleManager    *animationManager[Vector]
	rotation        rotation
	rotationManager *animationManager[rotation]
	colored         bool
	color           [4]uint8
	colorManager    *animationManager[colors]
	masked          bool
	mask            masker
	maskManager     *animationManager[masker]
}

func (ctr controller) Copy() controller {
	ctr.movementManager = ctr.movementManager.Copy()
	ctr.scaleManager = ctr.scaleManager.Copy()
	ctr.rotationManager = ctr.rotationManager.Copy()
	ctr.colorManager = ctr.colorManager.Copy()
	ctr.maskManager = ctr.maskManager.Copy()
	return ctr
}

func (ctr *controller) SetAlign(a Align) {
	ctr.align = a
}

func (ctr *controller) SetAnimation(a Animation) {
	ctr.animation = a

	if ctr.movementManager != nil {
		ctr.movementManager.SetAnimation(a)
	}

	if ctr.scaleManager != nil {
		ctr.scaleManager.SetAnimation(a)
	}

	if ctr.rotationManager != nil {
		ctr.rotationManager.SetAnimation(a)
	}

	if ctr.colorManager != nil {
		ctr.colorManager.SetAnimation(a)
	}
}

func (ctr *controller) GetAnimation() Animation {
	return ctr.animation
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

	if ctr.movementManager == nil {
		ctr.movementManager = &animationManager[Vector]{}
	}

	ctr.movementManager.AddAnimation(newAnimationValue(
		Vector{X: x - ctr.movement.X, Y: y - ctr.movement.Y},
		tick,
		len(replace) != 0 && replace[0],
		ctr.animation,
	))
}

func (ctr *controller) SetRotate(degree float64, replace ...bool) {
	if len(replace) != 0 && replace[0] {
		ctr.rotation = rotation(degree)
	} else {
		ctr.rotation = ctr.rotation.Add(rotation(degree))
	}
}

func (ctr *controller) SetRotating(degree float64, tick int, replace ...bool) {
	if tick <= 0 {
		return
	}

	if ctr.rotationManager == nil {
		ctr.rotationManager = &animationManager[rotation]{}
	}

	ctr.rotationManager.AddAnimation(newAnimationValue(
		rotation(degree).Sub(ctr.rotation),
		tick,
		len(replace) != 0 && replace[0],
		ctr.animation,
	))
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

	if ctr.scaleManager == nil {
		ctr.scaleManager = &animationManager[Vector]{}
	}

	ctr.scaleManager.AddAnimation(newAnimationValue(
		Vector{X: x - ctr.scale.X, Y: y - ctr.scale.Y},
		tick,
		len(replace) != 0 && replace[0],
		ctr.animation,
	))
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

	if ctr.colorManager == nil {
		ctr.colorManager = &animationManager[colors]{}
	}

	ctr.colorManager.AddAnimation(newAnimationValue(
		colors{r - ctr.color[0], g - ctr.color[1], b - ctr.color[2], a - ctr.color[3]},
		tick,
		true,
		ctr.animation,
	))
}

func (ctr *controller) SetMask(mk masker) {
	if !ctr.masked {
		ctr.mask = masker{0, 0, 1, 1}
		ctr.masked = true
	}

	ctr.mask = mk
}

func (ctr *controller) SetMasking(mk masker, tick int) {
	if tick <= 0 {
		return
	}

	if !ctr.masked {
		ctr.mask = masker{0, 0, 1, 1}
		ctr.masked = true
	}

	if ctr.maskManager == nil {
		ctr.maskManager = &animationManager[masker]{}
	}

	ctr.maskManager.AddAnimation(newAnimationValue(
		masker{
			X: mk.X - ctr.mask.X,
			Y: mk.Y - ctr.mask.Y,
			W: mk.W - ctr.mask.W,
			H: mk.H - ctr.mask.H,
		},
		tick,
		true,
		ctr.animation,
	))
}

func (ctr *controller) GetAlign() Align {
	return ctr.align
}

func (ctr *controller) GetMove() (x, y float64) {
	if ctr.movementManager != nil {
		ctr.movement = ctr.movementManager.GetAnimationResult(ctr.movement)
	}

	return ctr.movement.X, ctr.movement.Y
}

func (ctr *controller) GetDirection() Direction {
	return ctr.direction
}

func (ctr *controller) GetRotate() float64 {
	if ctr.rotationManager != nil {
		ctr.rotation = ctr.rotationManager.GetAnimationResult(ctr.rotation)
	}

	return ctr.rotation.Float64()
}

func (ctr *controller) GetScale() (x, y float64) {
	if !ctr.scaled {
		return 1, 1
	}

	if ctr.scaleManager != nil {
		ctr.scale = ctr.scaleManager.GetAnimationResult(ctr.scale)
	}

	return ctr.scale.X, ctr.scale.Y
}

func (ctr *controller) GetColor() (r, g, b, a uint8) {
	if !ctr.colored {
		return 255, 255, 255, 255
	}

	if ctr.colorManager != nil {
		ctr.color = ctr.colorManager.GetAnimationResult(ctr.color)
	}

	return ctr.color[0], ctr.color[1], ctr.color[2], ctr.color[3]
}

func (ctr *controller) GetMask() masker {
	if !ctr.masked {
		return masker{0, 0, 1, 1}
	}

	if ctr.maskManager != nil {
		ctr.mask = ctr.maskManager.GetAnimationResult(ctr.mask)
	}

	return ctr.mask
}

func (ctr *controller) GetBarycenter(parentMovement ...Vector) (float64, float64) {
	if len(parentMovement) == 0 {
		return ctr.movement.X, ctr.movement.Y
	}

	return ctr.movement.X + parentMovement[0].X, ctr.movement.Y + parentMovement[0].Y
}
