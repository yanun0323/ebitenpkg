package ebitenpkg

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

var _vertexRatio = alignHelper[[]Vector]{
	Center:         []Vector{{-0.5, -0.5}, {0.5, -0.5}, {0.5, 0.5}, {-0.5, 0.5}},
	TopCenter:      []Vector{{-0.5, 0}, {0.5, 0}, {0.5, 1}, {-0.5, 1}},
	BottomCenter:   []Vector{{-0.5, -1}, {0.5, -1}, {0.5, 0}, {-0.5, 0}},
	Leading:        []Vector{{0, -0.5}, {1, -0.5}, {1, 0.5}, {0, 0.5}},
	TopLeading:     []Vector{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
	BottomLeading:  []Vector{{0, -1}, {1, -1}, {1, 0}, {0, 0}},
	Trailing:       []Vector{{-1, -0.5}, {0, -0.5}, {0, 0.5}, {-1, 0.5}},
	TopTrailing:    []Vector{{-1, 0}, {0, 0}, {0, 1}, {-1, 1}},
	BottomTrailing: []Vector{{-1, -1}, {0, -1}, {0, 0}, {-1, 0}},
}

var _barycenter = alignHelper[func(w, h float64) (x, y float64)]{
	Center:         func(w, h float64) (x, y float64) { return 0.5 * w, 0.5 * h },
	TopCenter:      func(w, h float64) (x, y float64) { return 0.5 * w, 0 },
	BottomCenter:   func(w, h float64) (x, y float64) { return 0.5 * w, h },
	Leading:        func(w, h float64) (x, y float64) { return 0, 0.5 * h },
	TopLeading:     func(w, h float64) (x, y float64) { return 0, 0 },
	BottomLeading:  func(w, h float64) (x, y float64) { return 0, h },
	Trailing:       func(w, h float64) (x, y float64) { return w, 0.5 * h },
	TopTrailing:    func(w, h float64) (x, y float64) { return w, 0 },
	BottomTrailing: func(w, h float64) (x, y float64) { return w, h },
}

type Align int8

const (
	/*
		AlignCenter
			□□□
			□■□
			□□□
	*/
	AlignCenter Align = 0

	/*
		AlignTop
			□■□
			□□□
			□□□
	*/
	AlignTop Align = 1 << 1

	/*
		AlignTrailing
			□□□
			□□■
			□□□
	*/
	AlignTrailing Align = 1 << 2

	/*
		AlignBottom
			□□□
			□□□
			□■□
	*/
	AlignBottom Align = 1 << 3

	/*
		AlignLeading
			□□□
			■□□
			□□□
	*/
	AlignLeading Align = 1 << 4

	/*
		AlignTopLeading
			■□□
			□□□
			□□□
	*/
	AlignTopLeading Align = AlignTop | AlignLeading

	/*
		AlignBottomLeading
			□□□
			□□□
			■□□
	*/
	AlignBottomLeading Align = AlignBottom | AlignLeading
	/*
		AlignTopTrailing
			□□■
			□□□
			□□□
	*/
	AlignTopTrailing Align = AlignTop | AlignTrailing
	/*
		AlignBottomTrailing
			□□□
			□□□
			□□■
	*/
	AlignBottomTrailing Align = AlignBottom | AlignTrailing
)

func (a Align) String() string {
	return _string.Switch(a)
}

func (a Align) vertexRatio() []Vector {
	ratio := _vertexRatio.Switch(a)
	result := make([]Vector, len(ratio))
	copy(result, ratio)
	if len(result) != 4 {
		result = []Vector{{-0.5, -0.5}, {0.5, -0.5}, {0.5, 0.5}, {-0.5, 0.5}}
	}

	return result
}

func (a Align) barycenterOffset(w, h float64) (x, y float64) {
	return _barycenter.Switch(a)(w, h)
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
	f(AlignTop, h.TopCenter)
	f(AlignBottom, h.BottomCenter)
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
	case AlignTop:
		return h.TopCenter
	case AlignBottom:
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
