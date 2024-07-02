package ebitenpkg

type Align int

const (
	/*
		AlignCenter
			□□□
			□■□
			□□□
	*/
	AlignCenter Align = iota

	/*
		AlignTopCenter
			□■□
			□□□
			□□□
	*/
	AlignTopCenter

	/*
		AlignBottomCenter
			□□□
			□□□
			□■□
	*/
	AlignBottomCenter

	/*
		AlignLeading
			□□□
			■□□
			□□□
	*/
	AlignLeading

	/*
		AlignTopLeading
			■□□
			□□□
			□□□
	*/
	AlignTopLeading

	/*
		AlignBottomLeading
			□□□
			□□□
			■□□
	*/
	AlignBottomLeading
	/*
		AlignTrailing
			□□□
			□□■
			□□□
	*/
	AlignTrailing
	/*
		AlignTopTrailing
			□□■
			□□□
			□□□
	*/
	AlignTopTrailing
	/*
		AlignBottomTrailing
			□□□
			□□□
			□□■
	*/
	AlignBottomTrailing
)

var _string = alignHelper[string]{
	Center:         "Center",
	TopCenter:      "TopCenter",
	BottomCenter:   "BottomCenter",
	Leading:        "Leading",
	TopLeading:     "TopLeading",
	BottomLeading:  "BottomLeading",
	Trailing:       "Trailing",
	TopTrailing:    "TopTrailing",
	BottomTrailing: "BottomTrailing",
}

func (a Align) String() string {
	return _string.Switch(a)
}

type alignHelper[T any] struct {
	Center         T
	TopCenter      T
	BottomCenter   T
	Leading        T
	TopLeading     T
	BottomLeading  T
	Trailing       T
	TopTrailing    T
	BottomTrailing T
}

func (h alignHelper[T]) Execute(f func(Align, T)) {
	f(AlignCenter, h.Center)
	f(AlignTopCenter, h.TopCenter)
	f(AlignBottomCenter, h.BottomCenter)
	f(AlignLeading, h.Leading)
	f(AlignTopLeading, h.TopLeading)
	f(AlignBottomLeading, h.BottomLeading)
	f(AlignTrailing, h.Trailing)
	f(AlignTopTrailing, h.TopTrailing)
	f(AlignBottomTrailing, h.BottomTrailing)
}

func (h alignHelper[T]) Switch(a Align) T {
	switch a {
	case AlignCenter:
		return h.Center
	case AlignTopCenter:
		return h.TopCenter
	case AlignBottomCenter:
		return h.BottomCenter
	case AlignLeading:
		return h.Leading
	case AlignTopLeading:
		return h.TopLeading
	case AlignBottomLeading:
		return h.BottomLeading
	case AlignTrailing:
		return h.Trailing
	case AlignTopTrailing:
		return h.TopTrailing
	case AlignBottomTrailing:
		return h.BottomTrailing
	default:
		var zero T
		return zero
	}
}
