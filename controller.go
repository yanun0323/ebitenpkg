package ebitenpkg

import (
	sysimage "image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var radianToDegree = 180 / math.Pi

type vector struct {
	X, Y float64
}

type controller struct {
	reference vector
	center    vector
	movement  vector
	scale     vector
	rotation  float64
	align     Align
}

func NewController(referenceX, referenceY float64, a ...Align) Controller {
	c := newController(referenceX, referenceY, a...)
	return &c
}

func newController(referenceX, referenceY float64, a ...Align) controller {
	align := AlignCenter
	if len(a) != 0 {
		align = a[0]
	}

	return controller{
		reference: vector{X: referenceX, Y: referenceY},
		center:    vector{X: referenceX / 2, Y: referenceY / 2},
		movement:  vector{X: 0, Y: 0},
		scale:     vector{X: 1, Y: 1},
		align:     align,
	}
}

/*
	Controller
*/

func (f controller) Copy() Controller {
	return &f
}

func (f *controller) NewImage(img sysimage.Image) Image {
	return newImage(img, f)
}

func (f *controller) NewText(s string, size float64) Text {
	return newText(s, size, f)
}

/*
	embedController
*/

func (f *controller) Align(a Align) Controller {
	f.align = a
	return f
}

func (f *controller) Move(x, y float64, replace ...bool) Controller {
	if len(replace) != 0 && replace[0] {
		f.movement = vector{X: x, Y: y}
	} else {
		f.movement = vector{X: f.movement.X + x, Y: f.movement.Y + y}
	}

	return f
}

func (f *controller) Rotate(degree float64, replace ...bool) Controller {
	if len(replace) != 0 && replace[0] {
		f.rotation = degree
	} else {
		f.rotation += degree
	}

	return f
}

func (f *controller) Scale(x, y float64, replace ...bool) Controller {
	if len(replace) != 0 && replace[0] {
		f.scale = vector{X: x, Y: y}
	} else {
		f.scale = vector{X: f.scale.X * x, Y: f.scale.Y * y}
	}

	return f
}

func (f *controller) ReplaceController(ctr Controller) Controller {
	return ctr
}

func (f *controller) updateControllerReference() Controller {
	return f
}

func (f controller) Aligned() Align {
	return f.align
}

func (f controller) Moved() (x, y float64) {
	return f.movement.X, f.movement.Y
}

func (f controller) Rotated() float64 {
	return f.rotation
}

func (f controller) Scaled() (x, y float64) {
	return f.scale.X, f.scale.Y
}

func (f controller) DrawOption() *ebiten.DrawImageOptions {
	return f.drawOption()
}

func (f *controller) rotationCenter() vector {
	return f.movement
}

func (f *controller) vertexes() []vector {
	f.updateReferenceCenter()

	result := alignHelper[[]vector]{
		Center:         []vector{{-0.5, -0.5}, {0.5, -0.5}, {0.5, 0.5}, {-0.5, 0.5}},
		TopCenter:      []vector{{-0.5, 0}, {0.5, 0}, {0.5, 1}, {-0.5, 1}},
		BottomCenter:   []vector{{-0.5, -1}, {0.5, -1}, {0.5, 0}, {-0.5, 0}},
		Leading:        []vector{{0, -0.5}, {1, -0.5}, {1, 0.5}, {0, 0.5}},
		TopLeading:     []vector{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
		BottomLeading:  []vector{{0, -1}, {1, -1}, {1, 0}, {0, 0}},
		Trailing:       []vector{{-1, -0.5}, {0, -0.5}, {0, 0.5}, {-1, 0.5}},
		TopTrailing:    []vector{{-1, 0}, {0, 0}, {0, 1}, {-1, 1}},
		BottomTrailing: []vector{{-1, -1}, {0, -1}, {0, 0}, {-1, 0}},
	}.Switch(f.align)

	if len(result) != 4 {
		result = make([]vector, 4)
	}

	for i, v := range result {
		v.X *= f.reference.X
		v.Y *= f.reference.Y

		v = scaleVector(vector{}, v, f.scale)
		v = rotateVector(vector{}, v, f.rotation)
		v.X += f.movement.X
		v.Y += f.movement.Y
		result[i] = v
	}

	return result
}

func (f *controller) updateReference(x, y float64) {
	f.reference = vector{X: x, Y: y}
	f.updateReferenceCenter()
}

func (f *controller) drawOption() *ebiten.DrawImageOptions {
	f.updateReferenceCenter()
	oX, oY := f.alignOffset()
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-oX, -oY)
	opt.GeoM.Scale(f.scale.X, f.scale.Y)
	opt.GeoM.Rotate(f.rotation / radianToDegree)
	opt.GeoM.Translate(f.movement.X, f.movement.Y)
	return opt
}

/*
	private
*/

func (f *controller) updateReferenceCenter() {
	f.center = vector{X: f.reference.X / 2, Y: f.reference.Y / 2}
}

func (f controller) alignOffset() (x, y float64) {
	result := alignHelper[[2]float64]{
		Center:         [2]float64{f.center.X, f.center.Y},
		TopCenter:      [2]float64{f.center.X, 0},
		BottomCenter:   [2]float64{f.center.X, f.reference.Y},
		Leading:        [2]float64{0, f.center.Y},
		TopLeading:     [2]float64{0, 0},
		BottomLeading:  [2]float64{0, f.reference.Y},
		Trailing:       [2]float64{f.reference.X, f.center.Y},
		TopTrailing:    [2]float64{f.reference.X, 0},
		BottomTrailing: [2]float64{f.reference.X, f.reference.Y},
	}.Switch(f.align)

	return result[0], result[1]
}
