package ebitenpkg

// Value calculates the value ratio with the time passing
//
//   - timePassRatio: [0, 1]
//   - valueRatio: [0, 1]
type Animation func(timePassRatio float64) (valueRatio float64)

func AnimationDefault() Animation {
	return func(timePassRatio float64) float64 {
		return timePassRatio
	}
}

func AnimationNone() Animation {
	return func(float64) float64 {
		return 1
	}
}

func AnimationLinear() Animation {
	return func(timePassRatio float64) float64 {
		return timePassRatio
	}
}

func AnimationEaseInOut() Animation {
	return func(timePassRatio float64) float64 {
		if timePassRatio < 0.5 {
			return 2 * timePassRatio * timePassRatio
		}
		return -1 + (4-2*timePassRatio)*timePassRatio
	}
}

func AnimationEaseIn() Animation {
	return func(timePassRatio float64) float64 {
		// Classic quadratic ease‑in: starts slow, accelerates.
		return timePassRatio * timePassRatio
	}
}

func AnimationEaseOut() Animation {
	return func(timePassRatio float64) float64 {
		// Classic quadratic ease‑out: starts fast, decelerates.
		return 1 - (1-timePassRatio)*(1-timePassRatio)
	}
}
