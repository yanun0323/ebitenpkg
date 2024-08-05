package ebitenpkg

type Attachable interface {
	Aligned() Align
	Moved() (x, y float64)
	Scaled() (x, y float64)
	Rotated() (angle float64)
}
