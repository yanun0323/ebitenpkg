package ebitenpkg

type Vector struct {
	X, Y float64
}

func (v *Vector) Direction() Direction {
	var dir Direction
	switch {
	case v.X > 0:
		dir |= Right
	case v.X < 0:
		dir |= Left
	}

	switch {
	case v.Y > 0:
		dir |= Up
	case v.Y < 0:
		dir |= Down
	}

	return dir
}
