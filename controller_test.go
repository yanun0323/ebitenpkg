package ebitenpkg

import (
	"testing"
)

func TestControllerRotationCenter(t *testing.T) {
	testCases := []struct {
		desc     string
		ctr      func(Align) Controller
		expected alignHelper[vector]
	}{
		{
			"static",
			func(a Align) Controller {
				return NewController(2, 2, a)
			},
			alignHelper[vector]{
				Center:         vector{0, 0},
				TopCenter:      vector{0, 0},
				BottomCenter:   vector{0, 0},
				Leading:        vector{0, 0},
				TopLeading:     vector{0, 0},
				BottomLeading:  vector{0, 0},
				Trailing:       vector{0, 0},
				TopTrailing:    vector{0, 0},
				BottomTrailing: vector{0, 0},
			},
		},
		{
			"move",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(2, 4)
			},
			alignHelper[vector]{
				Center:         vector{2, 4},
				TopCenter:      vector{2, 4},
				BottomCenter:   vector{2, 4},
				Leading:        vector{2, 4},
				TopLeading:     vector{2, 4},
				BottomLeading:  vector{2, 4},
				Trailing:       vector{2, 4},
				TopTrailing:    vector{2, 4},
				BottomTrailing: vector{2, 4},
			},
		},
		{
			"move + scale",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(2, 4).Scale(5, 3)
			},
			alignHelper[vector]{
				Center:         vector{2, 4},
				TopCenter:      vector{2, 4},
				BottomCenter:   vector{2, 4},
				Leading:        vector{2, 4},
				TopLeading:     vector{2, 4},
				BottomLeading:  vector{2, 4},
				Trailing:       vector{2, 4},
				TopTrailing:    vector{2, 4},
				BottomTrailing: vector{2, 4},
			},
		},
		{
			"move + scale + rotate",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(2, 4).Scale(5, 3).Rotate(30)
			},
			alignHelper[vector]{
				Center:         vector{2, 4},
				TopCenter:      vector{2, 4},
				BottomCenter:   vector{2, 4},
				Leading:        vector{2, 4},
				TopLeading:     vector{2, 4},
				BottomLeading:  vector{2, 4},
				Trailing:       vector{2, 4},
				TopTrailing:    vector{2, 4},
				BottomTrailing: vector{2, 4},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.expected.Execute(func(a Align, expected vector) {
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
		expected alignHelper[[]vector]
	}{
		{
			"static",
			func(a Align) Controller {
				return NewController(2, 2, a)
			},
			alignHelper[[]vector]{
				Center:         []vector{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}},
				TopCenter:      []vector{{-1, 0}, {1, 0}, {1, 2}, {-1, 2}},
				BottomCenter:   []vector{{-1, -2}, {1, -2}, {1, 0}, {-1, 0}},
				Leading:        []vector{{0, -1}, {2, -1}, {2, 1}, {0, 1}},
				TopLeading:     []vector{{0, 0}, {2, 0}, {2, 2}, {0, 2}},
				BottomLeading:  []vector{{0, -2}, {2, -2}, {2, 0}, {0, 0}},
				Trailing:       []vector{{-2, -1}, {0, -1}, {0, 1}, {-2, 1}},
				TopTrailing:    []vector{{-2, 0}, {0, 0}, {0, 2}, {-2, 2}},
				BottomTrailing: []vector{{-2, -2}, {0, -2}, {0, 0}, {-2, 0}},
			},
		},
		{
			"move",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(3, 6)
			},
			alignHelper[[]vector]{
				Center:         []vector{{2, 5}, {4, 5}, {4, 7}, {2, 7}},
				TopCenter:      []vector{{2, 6}, {4, 6}, {4, 8}, {2, 8}},
				BottomCenter:   []vector{{2, 4}, {4, 4}, {4, 6}, {2, 6}},
				Leading:        []vector{{3, 5}, {5, 5}, {5, 7}, {3, 7}},
				TopLeading:     []vector{{3, 6}, {5, 6}, {5, 8}, {3, 8}},
				BottomLeading:  []vector{{3, 4}, {5, 4}, {5, 6}, {3, 6}},
				Trailing:       []vector{{1, 5}, {3, 5}, {3, 7}, {1, 7}},
				TopTrailing:    []vector{{1, 6}, {3, 6}, {3, 8}, {1, 8}},
				BottomTrailing: []vector{{1, 4}, {3, 4}, {3, 6}, {1, 6}},
			},
		},
		{
			"move + scale",
			func(a Align) Controller {
				return NewController(2, 2, a).Move(8, 6).Scale(3, 2)
			},
			alignHelper[[]vector]{
				Center:         []vector{{5, 4}, {11, 4}, {11, 8}, {5, 8}},
				TopCenter:      []vector{{5, 6}, {11, 6}, {11, 10}, {5, 10}},
				BottomCenter:   []vector{{5, 2}, {11, 2}, {11, 6}, {5, 6}},
				Leading:        []vector{{8, 4}, {14, 4}, {14, 8}, {8, 8}},
				TopLeading:     []vector{{8, 6}, {14, 6}, {14, 10}, {8, 10}},
				BottomLeading:  []vector{{8, 2}, {14, 2}, {14, 6}, {8, 6}},
				Trailing:       []vector{{2, 4}, {8, 4}, {8, 8}, {2, 8}},
				TopTrailing:    []vector{{2, 6}, {8, 6}, {8, 10}, {2, 10}},
				BottomTrailing: []vector{{2, 2}, {8, 2}, {8, 6}, {2, 6}},
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
			tc.expected.Execute(func(a Align, expected []vector) {
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
