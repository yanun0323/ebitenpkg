package ebitenpkg

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var _radianToDegree = 180 / math.Pi

type vector struct {
	X, Y float64
}

type DrawOption struct {
	*drawOption
}

func NewDrawOption(referenceX, referenceY float64, a ...Align) *DrawOption {
	return &DrawOption{
		drawOption: newDrawOption(referenceX, referenceY, a...),
	}
}

func (f DrawOption) WithReference(x, y float64) *DrawOption {
	f.drawOption = f.withReference(x, y)
	return &f
}

func (f *DrawOption) UpdateReference(x, y float64) {
	f.updateReference(x, y)
}

func (f DrawOption) Text(text string, size float64, a ...Align) *Text {
	return NewText(text, size, a...).WithOption(f)
}

func (f DrawOption) Image(img *ebiten.Image, a ...Align) *Image {
	return NewImage(img, a...).WithOption(f)
}

type drawOption struct {
	reference vector
	center    vector
	movement  vector
	scale     vector
	rotation  float64
	align     Align
}

func newDrawOption(referenceX, referenceY float64, a ...Align) *drawOption {
	align := AlignCenter
	if len(a) != 0 {
		align = a[0]
	}

	return &drawOption{
		reference: vector{X: referenceX, Y: referenceY},
		center:    vector{X: referenceX / 2, Y: referenceY / 2},
		movement:  vector{X: 1, Y: 1},
		scale:     vector{X: 1, Y: 1},
		align:     align,
	}
}

func (f drawOption) Option() *ebiten.DrawImageOptions {
	return f.option()
}

func (f drawOption) Movement() (x, y float64) {
	return f.movement.X, f.movement.Y
}

func (f drawOption) ScaleRatio() (x, y float64) {
	return f.scale.X, f.scale.Y
}

func (f drawOption) Rotation() float64 {
	return f.rotation
}

func (f drawOption) Alignment() Align {
	return f.align
}

func (f drawOption) copy() *drawOption {
	ff := f
	return &ff
}

func (f drawOption) withMovement(x, y float64, replace ...bool) *drawOption {
	f.Move(x, y, replace...)
	return &f
}

func (f drawOption) withScaleRatio(x, y float64, replace ...bool) *drawOption {
	f.Scale(x, y, replace...)
	return &f
}

func (f drawOption) withRotation(degree float64, replace ...bool) *drawOption {
	f.Rotate(degree, replace...)
	return &f
}

func (f drawOption) withAlignment(a Align) *drawOption {
	f.Align(a)
	return &f
}

func (f drawOption) withReference(x, y float64) *drawOption {
	f.reference = vector{X: x, Y: y}
	f.recalculateCenter()
	return &f
}

func (f *drawOption) Move(x, y float64, replace ...bool) {
	if len(replace) != 0 && replace[0] {
		f.movement = vector{X: x, Y: y}
	} else {
		f.movement = vector{X: f.movement.X + x, Y: f.movement.Y + y}
	}
}

func (f *drawOption) Scale(x, y float64, replace ...bool) {
	if len(replace) != 0 && replace[0] {
		f.scale = vector{X: x, Y: y}
	} else {
		f.scale = vector{X: f.scale.X * x, Y: f.scale.Y * y}
	}
}

func (f *drawOption) Rotate(degree float64, replace ...bool) {
	if len(replace) != 0 && replace[0] {
		f.rotation = degree
	} else {
		f.rotation += degree
	}
}

func (f *drawOption) Align(a Align) {
	f.align = a
}

func (f *drawOption) updateReference(x, y float64) {
	f.reference = vector{X: x, Y: y}
	f.recalculateCenter()
}

func (f *drawOption) recalculateCenter() {
	f.center = vector{X: f.reference.X / 2, Y: f.reference.Y / 2}
}

func (f drawOption) alignOffset() (x, y float64) {
	switch f.align {
	case AlignCenter:
		return f.center.X, f.center.Y
	case AlignTopCenter:
		return f.center.X, 0
	case AlignBottomCenter:
		return f.center.X, f.reference.Y

	case AlignLeading:
		return 0, f.center.Y
	case AlignTopLeading:
		return 0, 0
	case AlignBottomLeading:
		return 0, f.reference.Y

	case AlignTrailing:
		return f.reference.X, f.center.Y
	case AlignTopTrailing:
		return f.reference.X, 0
	case AlignBottomTrailing:
		return f.reference.X, f.reference.Y
	default:
		return f.center.X, f.center.Y
	}
}

func (f *drawOption) option() *ebiten.DrawImageOptions {
	f.recalculateCenter()
	oX, oY := f.alignOffset()
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-oX, -oY)
	opt.GeoM.Rotate(f.rotation / _radianToDegree)
	opt.GeoM.Scale(f.scale.X, f.scale.Y)
	opt.GeoM.Translate(f.movement.X, f.movement.Y)
	return opt
}

func (f *drawOption) debugOption(borderWidth ...int) *ebiten.DrawImageOptions {
	b := float64(_defaultDebugBorderWidth)
	if len(borderWidth) != 0 && borderWidth[0] >= 0 {
		b = float64(borderWidth[0])
	}

	f.recalculateCenter()
	oX, oY := f.alignOffset()
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-oX-b, -oY-b)
	opt.GeoM.Scale(f.scale.X, f.scale.Y)
	opt.GeoM.Rotate(f.rotation / _radianToDegree)
	opt.GeoM.Translate(f.movement.X, f.movement.Y)
	return opt
}
