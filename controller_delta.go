package ebitenpkg

import (
	"sync"
)

type deltaValue[T any] interface {
	Sub(T) T
	Add(T) T
	Ratio(float64) T
}

type controllerDelta[T deltaValue[T]] struct {
	mu         sync.RWMutex
	tick       int
	startTick  int
	targetTick int
	lastTick   int
	animation  Animation

	param          T
	isParamReplace bool
}

func newControllerDelta[T deltaValue[T]](param T, tick int, replace bool, animation Animation) (add *controllerDelta[T], isReplace bool) {
	var (
		current = CurrentGameTime()
		a       = animation
	)

	if a == nil {
		a = AnimationDefault()
	}

	return &controllerDelta[T]{
		tick:           tick,
		startTick:      current,
		targetTick:     current + tick,
		lastTick:       current,
		animation:      a,
		param:          param,
		isParamReplace: replace,
	}, replace
}

func (c *controllerDelta[T]) IsComplete() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lastTick >= c.targetTick
}

func (c *controllerDelta[T]) CalculateResult(currentTick int, value T) T {
	c.mu.Lock()
	defer c.mu.Unlock()

	if currentTick > c.targetTick {
		if c.isParamReplace {
			return c.param
		}

		return value
	}

	delta := c.param
	if c.isParamReplace {
		delta = c.param.Sub(value)
	}

	base := float64(c.targetTick) - float64(c.startTick)
	lastTimePassRatio := (float64(c.lastTick) - float64(c.startTick)) / base
	lastRatio := c.animation(lastTimePassRatio)
	if lastRatio+_floatFix >= 1 {
		lastRatio = 1
	}
	lastDelta := delta.Ratio(lastRatio)

	currTimePassRatio := (float64(currentTick+1) - float64(c.startTick)) / base
	currRatio := c.animation(currTimePassRatio)
	if currRatio+_floatFix >= 1 {
		currRatio = 1
	}
	currDelta := delta.Ratio(currRatio)

	delta = currDelta.Sub(lastDelta)
	c.lastTick = currentTick

	return value.Add(delta)
}
