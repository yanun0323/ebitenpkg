package ebitenpkg

import (
	"sync"
)

type controllerDelta struct {
	mu         sync.RWMutex
	tick       int
	startTick  int
	targetTick int
	lastTick   int
	animation  Animation

	param          Vector
	isParamReplace bool
}

func newControllerDelta(x, y float64, tick int, replace bool, animation Animation) (add *controllerDelta, isReplace bool) {
	var (
		current = CurrentGameTime()
		a       = animation
	)

	if a == nil {
		a = AnimationDefault()
	}

	return &controllerDelta{
		tick:           tick,
		startTick:      current,
		targetTick:     current + tick,
		lastTick:       current,
		animation:      a,
		param:          Vector{x, y},
		isParamReplace: replace,
	}, replace
}

func (c *controllerDelta) IsComplete() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lastTick >= c.targetTick
}

func (c *controllerDelta) CalculateResult(currentTick int, value Vector) Vector {
	c.mu.Lock()
	defer c.mu.Unlock()

	if currentTick >= c.targetTick {
		return value
	}

	delta := c.param
	if c.isParamReplace {
		delta = Vector{
			X: c.param.X - value.X,
			Y: c.param.Y - value.Y,
		}
	}

	lastTimePassRatio := (float64(c.lastTick) - float64(c.startTick)) / float64(c.tick)
	lastRatio := c.animation(lastTimePassRatio)

	currTimePassRatio := (float64(currentTick) - float64(c.startTick)) / float64(c.tick)
	currRatio := c.animation(currTimePassRatio)

	ratio := currRatio - lastRatio

	c.lastTick = currentTick

	return Vector{
		X: value.X + delta.X*ratio,
		Y: value.Y + delta.Y*ratio,
	}
}
