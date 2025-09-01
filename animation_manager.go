package ebitenpkg

import (
	"sync"
)

// AnimationManager 管理所有動畫類型的動畫
type animationManager[T animatableValue[T]] struct {
	animations []*animatedValue[T]
	mu         sync.RWMutex
	lastTick   int
}

// animatedValue 代表一個具體的動畫值
type animatedValue[T animatableValue[T]] struct {
	lastValue  T
	deltaValue T
	duration   int // Tick
	startTime  int // Tick
	animation  Animation
	isComplete bool
	replace    bool // 是否替換舊動畫
}

func newAnimationValue[T animatableValue[T]](delta T, duration int, replace bool, animation Animation) *animatedValue[T] {
	return &animatedValue[T]{
		deltaValue: delta,
		duration:   duration,
		startTime:  CurrentGameTime(),
		animation:  animation,
		isComplete: false,
		replace:    replace,
	}
}

func (am *animationManager[T]) SetAnimation(a Animation) {
	am.mu.Lock()
	defer am.mu.Unlock()

	for _, anim := range am.animations {
		anim.animation = a
	}
}

// 新增動畫
func (am *animationManager[T]) AddAnimation(anim *animatedValue[T]) {
	am.mu.Lock()
	defer am.mu.Unlock()

	// 如果是替換類型，先清理舊的動畫
	if anim.replace {
		am.animations = nil
	}

	am.animations = append(am.animations, anim)
}

// 獲取指定類型的活躍動畫
func (am *animationManager[T]) GetActiveAnimations() []*animatedValue[T] {
	am.mu.RLock()
	defer am.mu.RUnlock()

	if len(am.animations) == 0 {
		return nil
	}

	active := make([]*animatedValue[T], 0)
	for _, anim := range am.animations {
		if !anim.isComplete {
			active = append(active, anim)
		}
	}

	return active
}

// 獲取指定類型的動畫結果
func (am *animationManager[T]) GetAnimationResult(currentValue T) T {
	am.mu.RLock()
	defer am.mu.RUnlock()

	if len(am.animations) == 0 {
		return currentValue
	}

	active := am.animations
	if len(active) == 0 {
		return currentValue
	}

	currentTick := CurrentGameTime()
	if currentTick <= am.lastTick {
		return currentValue
	}

	// 簡單的疊加計算（根據動畫類型）
	result := currentValue
	for _, anim := range active {
		if !anim.isComplete {
			// 計算當前動畫結果並疊加到結果上
			result = am.calculateAnimationResult(anim, currentTick, result)
		}
	}

	am.lastTick = currentTick

	return result
}

// 計算單個動畫的結果
func (*animationManager[T]) calculateAnimationResult(anim *animatedValue[T], currentTick int, currentValue T) T {
	// 計算進度
	elapsed := currentTick - anim.startTime
	if elapsed >= anim.duration {
		anim.isComplete = true
		return currentValue
	}

	progress := float64(elapsed) / float64(anim.duration)

	animation := anim.animation
	if animation == nil {
		animation = AnimationDefault()
	}

	ratio := animation(progress)

	lastValue := anim.lastValue
	nextValue := anim.deltaValue.Ratio(ratio)
	anim.lastValue = nextValue
	nextValueDelta := nextValue.Sub(lastValue)

	return currentValue.Add(nextValueDelta)
}
