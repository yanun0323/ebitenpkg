package ebitenpkg

type Direction int8

const (
	DirectionUp    Direction = 1 << 1
	DirectionRight Direction = 1 << 2
	DirectionDown  Direction = 1 << 3
	DirectionLeft  Direction = 1 << 4
)

func newDirection(x, y float64) Direction {
	if x > 0 {
		if y > 0 {
			return DirectionUp
		} else {
			return DirectionDown
		}
	} else {
		if y > 0 {
			return DirectionRight
		} else {
			return DirectionLeft
		}
	}
}

func newDirectionFrom(x, y, x2, y2 float64) Direction {
	return newDirection(x2-x, y2-y)
}
