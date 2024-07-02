package ebitenpkg

import (
	"testing"
)

func TestControllerRotationCenter(t *testing.T) {
	testCases := []struct {
		desc     string
		ctr      func(Align) Controller
		expected alignHelper[Vector]
	}{
		{
			"static",
			func(a Align) Controller {
				return NewController(2, 2, a)
			},
			alignHelper[Vector]{
				Center:         Vector{0, 0},
				TopCenter:      Vector{0, 0},
				BottomCenter:   Vector{0, 0},
				Leading:        Vector{0, 0},
				TopLeading:     Vector{0, 0},
				BottomLeading:  Vector{0, 0},
				Trailing:       Vector{0, 0},
				TopTrailing:    Vector{0, 0},
				BottomTrailing: Vector{0, 0},
			},
		},
		{
			"move",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(2, 4)
			},
			alignHelper[Vector]{
				Center:         Vector{2, 4},
				TopCenter:      Vector{2, 4},
				BottomCenter:   Vector{2, 4},
				Leading:        Vector{2, 4},
				TopLeading:     Vector{2, 4},
				BottomLeading:  Vector{2, 4},
				Trailing:       Vector{2, 4},
				TopTrailing:    Vector{2, 4},
				BottomTrailing: Vector{2, 4},
			},
		},
		{
			"move + scale",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(2, 4).Scale(5, 3)
			},
			alignHelper[Vector]{
				Center:         Vector{2, 4},
				TopCenter:      Vector{2, 4},
				BottomCenter:   Vector{2, 4},
				Leading:        Vector{2, 4},
				TopLeading:     Vector{2, 4},
				BottomLeading:  Vector{2, 4},
				Trailing:       Vector{2, 4},
				TopTrailing:    Vector{2, 4},
				BottomTrailing: Vector{2, 4},
			},
		},
		{
			"move + scale + rotate",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(2, 4).Scale(5, 3).Rotate(30)
			},
			alignHelper[Vector]{
				Center:         Vector{2, 4},
				TopCenter:      Vector{2, 4},
				BottomCenter:   Vector{2, 4},
				Leading:        Vector{2, 4},
				TopLeading:     Vector{2, 4},
				BottomLeading:  Vector{2, 4},
				Trailing:       Vector{2, 4},
				TopTrailing:    Vector{2, 4},
				BottomTrailing: Vector{2, 4},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.expected.Execute(func(a Align, expected Vector) {
				ctr := tc.ctr(a)

				rc := ctr.rotationCenter()
				if rc.X != expected.X || rc.Y != expected.Y {
					t.Fatalf("[%s] expected %v, but got %v", a.String(), expected, rc)
				}
			})
		})
	}
}

func TestControllerVertexes(t *testing.T) {
	testCases := []struct {
		desc     string
		ctr      func(Align) Controller
		expected alignHelper[[]Vector]
	}{
		{
			"static",
			func(a Align) Controller {
				return NewController(2, 2, a)
			},
			alignHelper[[]Vector]{
				Center:         []Vector{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}},
				TopCenter:      []Vector{{-1, 0}, {1, 0}, {1, 2}, {-1, 2}},
				BottomCenter:   []Vector{{-1, -2}, {1, -2}, {1, 0}, {-1, 0}},
				Leading:        []Vector{{0, -1}, {2, -1}, {2, 1}, {0, 1}},
				TopLeading:     []Vector{{0, 0}, {2, 0}, {2, 2}, {0, 2}},
				BottomLeading:  []Vector{{0, -2}, {2, -2}, {2, 0}, {0, 0}},
				Trailing:       []Vector{{-2, -1}, {0, -1}, {0, 1}, {-2, 1}},
				TopTrailing:    []Vector{{-2, 0}, {0, 0}, {0, 2}, {-2, 2}},
				BottomTrailing: []Vector{{-2, -2}, {0, -2}, {0, 0}, {-2, 0}},
			},
		},
		{
			"move",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(3, 6)
			},
			alignHelper[[]Vector]{
				Center:         []Vector{{2, 5}, {4, 5}, {4, 7}, {2, 7}},
				TopCenter:      []Vector{{2, 6}, {4, 6}, {4, 8}, {2, 8}},
				BottomCenter:   []Vector{{2, 4}, {4, 4}, {4, 6}, {2, 6}},
				Leading:        []Vector{{3, 5}, {5, 5}, {5, 7}, {3, 7}},
				TopLeading:     []Vector{{3, 6}, {5, 6}, {5, 8}, {3, 8}},
				BottomLeading:  []Vector{{3, 4}, {5, 4}, {5, 6}, {3, 6}},
				Trailing:       []Vector{{1, 5}, {3, 5}, {3, 7}, {1, 7}},
				TopTrailing:    []Vector{{1, 6}, {3, 6}, {3, 8}, {1, 8}},
				BottomTrailing: []Vector{{1, 4}, {3, 4}, {3, 6}, {1, 6}},
			},
		},
		{
			"move + scale",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(8, 6).Scale(3, 2)
			},
			alignHelper[[]Vector]{
				Center:         []Vector{{5, 4}, {11, 4}, {11, 8}, {5, 8}},
				TopCenter:      []Vector{{5, 6}, {11, 6}, {11, 10}, {5, 10}},
				BottomCenter:   []Vector{{5, 2}, {11, 2}, {11, 6}, {5, 6}},
				Leading:        []Vector{{8, 4}, {14, 4}, {14, 8}, {8, 8}},
				TopLeading:     []Vector{{8, 6}, {14, 6}, {14, 10}, {8, 10}},
				BottomLeading:  []Vector{{8, 2}, {14, 2}, {14, 6}, {8, 6}},
				Trailing:       []Vector{{2, 4}, {8, 4}, {8, 8}, {2, 8}},
				TopTrailing:    []Vector{{2, 6}, {8, 6}, {8, 10}, {2, 10}},
				BottomTrailing: []Vector{{2, 2}, {8, 2}, {8, 6}, {2, 6}},
			},
		},
		// {
		// 	"move + scale + rotate",
		// 	func(a Align) Controller {
		// 		return NewController(2, 2, a).Move(8, 6).Scale(3, 2).Rotate(90)
		// 	},
		// 	alignHelper[[]vector]{
		// 		Center:         []vector{{5, 4}, {11, 4}, {5, 8}, {11, 8}},
		// 		TopCenter:      []vector{{5, 6}, {11, 6}, {5, 10}, {11, 10}},
		// 		BottomCenter:   []vector{{5, 2}, {11, 2}, {5, 6}, {11, 6}},
		// 		Leading:        []vector{{8, 4}, {14, 4}, {8, 8}, {14, 8}},
		// 		TopLeading:     []vector{{8, 6}, {14, 6}, {8, 10}, {14, 10}},
		// 		BottomLeading:  []vector{{8, 2}, {14, 2}, {8, 6}, {14, 6}},
		// 		Trailing:       []vector{{2, 4}, {8, 4}, {2, 8}, {8, 8}},
		// 		TopTrailing:    []vector{{2, 6}, {8, 6}, {2, 10}, {8, 10}},
		// 		BottomTrailing: []vector{{2, 2}, {8, 2}, {2, 6}, {8, 6}},
		// 	},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.expected.Execute(func(a Align, expected []Vector) {
				vs := tc.ctr(a).vertexes()
				if len(vs) != len(expected) {
					t.Fatalf("expected len %d, but got %d", len(expected), len(vs))
				}

				for i, e := range expected {
					if vs[i].X != e.X || vs[i].Y != e.Y {
						t.Fatalf("\n[%s] expected at %d should be \n%v, but got \n%v", a.String(), i, expected, vs)
					}
				}
			})
		})
	}
}
