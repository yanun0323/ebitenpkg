package ebitenpkg

type Direction int8

const (
	DirectionNone      Direction = 0
	DirectionUp        Direction = 1 << 1
	DirectionRight     Direction = 1 << 2
	DirectionDown      Direction = 1 << 3
	DirectionLeft      Direction = 1 << 4
	DirectionUpRight   Direction = DirectionUp | DirectionRight
	DirectionDownRight Direction = DirectionDown | DirectionRight
	DirectionDownLeft  Direction = DirectionDown | DirectionLeft
	DirectionUpLeft    Direction = DirectionUp | DirectionLeft
)

func newDirection(x, y float64) Direction {
	var d Direction

	switch {
	case x > 0:
		d |= DirectionRight
	case x < 0:
		d |= DirectionLeft
	}

	switch {
	case y > 0:
		d |= DirectionDown
	case y < 0:
		d |= DirectionUp
	}

	return d
}

func newDirectionFrom(x, y, x2, y2 float64) Direction {
	return newDirection(x2-x, y2-y)
}
