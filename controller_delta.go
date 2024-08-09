package ebitenpkg

import (
	"sync"
)

type controllerDelta struct {
	mu         sync.RWMutex
	tick       int
	targetTick int
	lastTick   int
	replace    bool   /* true: result, false: delta */
	tickDelta  Vector /* for not replace */
	tickResult Vector /* for replace */
}

func newControllerDelta(x, y float64, tick int, rp bool) (add *controllerDelta, replace bool) {
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

	return &controllerDelta{
		tick:       tick,
		targetTick: current + tick,
		lastTick:   current,
		replace:    rp,
		tickDelta:  delta,
		tickResult: result,
	}, rp
}

func (c *controllerDelta) IsComplete() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lastTick >= c.targetTick
}

func (c *controllerDelta) CalculateResult(currentTick int, value Vector, isScale bool) Vector {
	offset := c.getTickerOffset(currentTick)
	c.mu.RLock()
	defer c.mu.RUnlock()

	if offset <= _floatFix {
		return value
	}

	if c.replace {
		result := Vector{
			X: value.X + ((c.tickResult.X - value.X) / offset),
			Y: value.Y + ((c.tickResult.Y - value.Y) / offset),
		}

		return result
	}

	result := Vector{
		X: value.X + c.tickDelta.X*offset,
		Y: value.Y + c.tickDelta.Y*offset,
	}

	return result
}

func (c *controllerDelta) getTickerOffset(currentTick int) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.replace {
		result := float64(c.targetTick - c.lastTick)
		c.lastTick = currentTick
		return result
	}

	var result float64
	if currentTick >= c.targetTick {
		result = float64(c.targetTick - c.lastTick)
	} else {
		result = float64(currentTick - c.lastTick)
	}

	c.lastTick = currentTick
	return result
}
