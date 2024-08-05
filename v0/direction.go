package ebitenpkg

type Direction int

const (
	Left  Direction = 1 << 0
	Right Direction = 1 << 1
	Up    Direction = 1 << 2
	Down  Direction = 1 << 3
)

func newDirection[N int | float64](x, y N) Direction {
	return newDirectionFrom(0, 0, x, y)
}

func newDirectionFrom[N int | float64](oldX, oldY, newX, newY N) Direction {
	x := newX - oldX
	y := newY - oldY

	var d Direction
	switch {
	case x > 0:
		d |= Right
	case x < 0:
		d |= Left
	case y > 0:
		d |= Down
	case y < 0:
		d |= Up
	}

	return d
}
