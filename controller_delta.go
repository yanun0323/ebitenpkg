package ebitenpkg

type controllerDelta struct {
	tick       int
	targetTick int
	lastTick   int
	replace    bool   /* true: result, false: delta */
	tickDelta  Vector /* for not replace */
	tickResult Vector /* for replace */
}

func newControllerDelta(x, y float64, tick int, rp bool) (add controllerDelta, replace bool) {
	var (
		current = CurrentGameTime()
		delta   Vector
		result  Vector
	)

	if rp {
		result = Vector{X: x, Y: y}
	} else {
		delta = Vector{X: x / float64(tick), Y: y / float64(tick)}
	}

	return controllerDelta{
		tick:       tick,
		targetTick: current + tick,
		lastTick:   current,
		replace:    rp,
		tickDelta:  delta,
		tickResult: result,
	}, rp
}

func (c *controllerDelta) IsComplete() bool {
	return c.lastTick >= c.targetTick
}

func (c *controllerDelta) CalculateResult(currentTick int, value Vector, isScale bool) Vector {
	offset := c.getTickerOffset(currentTick)
	if offset <= 0 {
		return value
	}

	if c.replace {
		result := Vector{
			X: value.X + ((c.tickResult.X - value.X) / offset),
			Y: value.Y + ((c.tickResult.Y - value.Y) / offset),
		}

		return result
	}

	if isScale {
		return Vector{
			X: value.X + c.tickDelta.X*offset,
			Y: value.Y + c.tickDelta.Y*offset,
		}
	}

	return Vector{
		X: value.X + c.tickDelta.X*offset,
		Y: value.Y + c.tickDelta.Y*offset,
	}
}

func (c *controllerDelta) getTickerOffset(currentTick int) float64 {
	defer func() {
		c.lastTick = currentTick
	}()

	if c.replace {
		return float64(c.targetTick - c.lastTick)
	}

	if currentTick >= c.targetTick {
		return float64(c.targetTick - c.lastTick)
	} else {
		return float64(currentTick - c.lastTick)
	}
}
