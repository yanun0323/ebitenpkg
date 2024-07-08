package ebitenpkg

import (
	"math"
	"testing"
)

func TestMathIsInside(t *testing.T) {
	testCases := []struct {
		desc     string
		area     []Vector
		p        Vector
		expected bool
	}{
		{
			"good",
			[]Vector{{0, 0}, {4, 0}, {4, 4}, {0, 4}},
			Vector{2, 2},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := isInside(tc.area, tc.p)
			if result != tc.expected {
				t.Fatalf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestPolygonArea(t *testing.T) {
	testCases := []struct {
		desc     string
		area     []Vector
		expected float64
	}{
		{
			"square",
			[]Vector{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
			1.0,
		},
		{
			"square exclude zero",
			[]Vector{{-1, -1}, {1, -1}, {1, 1}, {-1, -1}},
			2.0,
		},
		{
			"rectangle",
			[]Vector{{0, 0}, {8, 0}, {8, 4}, {0, 4}},
			32.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := polygonArea(tc.area[0], tc.area[1], tc.area[2], tc.area[3])
			if math.Abs(result-tc.expected) >= 1 {
				t.Fatalf("expected %.2f, but got %.2f", tc.expected, result)
			}
		})
	}
}

func TestTriangleArea(t *testing.T) {
	testCases := []struct {
		desc     string
		area     []Vector
		expected float64
	}{
		{
			"right-triangle",
			[]Vector{{0, 0}, {4, 0}, {0, 4}},
			8.0,
		},
		{
			"isosceles-triangle",
			[]Vector{{0, 0}, {8, 0}, {4, 4}},
			16.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := triangleArea(tc.area[0], tc.area[1], tc.area[2])
			if math.Abs(result-tc.expected) >= 1 {
				t.Fatalf("expected %.2f, but got %.2f", tc.expected, result)
			}
		})
	}
}

func TestGetVertexes(t *testing.T) {
	testCases := []struct {
		desc string
		alignHelper[[]Vector]
	}{
		{
			"static",
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
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.alignHelper.Execute(func(a Align, v []Vector) {
				ctr := newController(a)
				vs := getVertexes(2, 2, ctr)
				if len(vs) != len(v) {
					t.Fatalf("expected len %d, but got %d", len(v), len(vs))
				}

				for i, e := range v {
					if vs[i].X != e.X || vs[i].Y != e.Y {
						t.Fatalf("\n[%s] expected at %d should be \n%v, but got \n%v", a.String(), i, v, vs)
					}
				}
			})
		})
	}
}
