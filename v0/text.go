package ebitenpkg

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	ebitenfonts "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	gofont "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	_dpi = 72
)

type text struct {
	ctr             controller
	debugImageCache debugCache
	s               string
	size            float64
	face            ebitentext.Face
	color           color.Color
	lineSpacing     float64
}

func NewText(s string, size float64, a Align) Text {
	return &text{
		ctr:   newController(a),
		s:     s,
		size:  size,
		face:  text{}.newFace(size),
		color: color.Black,
	}
}

/*
	Drawable
*/

func (t text) Draw(screen *ebiten.Image, debug ...color.Color) {
	w, h := t.Bounds()
	opt := t.DrawOption()

	ebitentext.Draw(screen, t.s, t.face, &ebitentext.DrawOptions{
		DrawImageOptions: *opt,
		LayoutOptions:    ebitentext.LayoutOptions{LineSpacing: t.lineSpacing},
	})

	if len(debug) != 0 && debug[0] != nil {
		debugImage := t.debugImageCache.Image(int(w), int(h), debug[0])
		screen.DrawImage(debugImage, opt)
	}
}

/*
	Controllable
*/

func (t *text) Align(a Align) Text {
	t.ctr.Align(a)
	return t
}

func (t *text) Move(x, y float64, replace ...bool) Text {
	t.ctr.Move(x, y, replace...)
	return t
}

func (t *text) Rotate(degree float64, replace ...bool) Text {
	t.ctr.Rotate(degree, replace...)
	return t
}

func (t *text) Scale(x, y float64, replace ...bool) Text {
	t.ctr.Scale(x, y, replace...)
	return t
}

func (t text) Aligned() Align {
	return t.ctr.Aligned()
}

func (t text) Moved() (x, y float64) {
	return t.ctr.Moved()
}

func (t text) Rotated() float64 {
	return t.ctr.Rotated()
}

func (t text) Scaled() (x, y float64) {
	return t.ctr.Scaled()
}

func (t text) DrawOption() *ebiten.DrawImageOptions {
	w, h := t.Bounds()
	return getDrawOption(w, h, t.ctr)
}

func (t text) Bounds() (w, h float64) {
	return ebitentext.Measure(t.s, t.face, t.lineSpacing)
}

func (t text) Barycenter() (x, y float64) {
	return t.ctr.Moved()
}

/*
	Text
*/

func (t text) Copy() Text {
	t.face = t.newFace(t.Size())
	t.debugImageCache = t.debugImageCache.Copy()

	return &t
}

func (t *text) SetColor(c color.Color) Text {
	t.color = c
	return t
}

func (t *text) SetLineSpacing(l float64) Text {
	t.lineSpacing = l
	return t
}

func (t *text) SetText(text string) Text {
	t.s = text
	return t
}

func (t *text) SetSize(size float64) Text {
	t.size = size
	t.face = t.newFace(size)
	return t
}

func (t text) Color() color.Color {
	return t.color
}

func (t text) LineSpacing() float64 {
	return t.lineSpacing
}

func (t text) Text() string {
	return t.s
}

func (t text) Size() float64 {
	return t.size
}

/*
	private
*/

func (text) newFace(size float64) ebitentext.Face {
	opt := &opentype.FaceOptions{
		Size:    size,
		DPI:     _dpi,
		Hinting: gofont.HintingNone,
	}

	ff, err := opentype.Parse(ebitenfonts.MPlus1pRegular_ttf)
	if err != nil {
		return ebitentext.NewGoXFace(nil)
	}

	ft, err := opentype.NewFace(ff, opt)
	if err != nil {
		return ebitentext.NewGoXFace(nil)
	}

	return ebitentext.NewGoXFace(ft)
}
