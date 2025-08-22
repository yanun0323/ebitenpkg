package ebitenpkg

type Vector struct {
	X, Y float64
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector) Ratio(r float64) Vector {
	return Vector{
		X: v.X * r,
		Y: v.Y * r,
	}
}
