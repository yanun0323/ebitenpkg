package ebitenpkg

import (
	"image/color"

	ebitenfonts "github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	gofont "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	_dpi = 72
)

type text struct {
	Controller

	s           string
	size        float64
	face        ebitentext.Face
	color       color.Color
	lineSpacing float64

	debugImg debugImage
}

func NewText(s string, size float64, a ...Align) Text {
	ctr := NewController(0, 0, a...)
	return (&text{
		s:          s,
		size:       size,
		face:       text{}.newFace(size),
		Controller: ctr,
		debugImg:   newDebugImage(ctr),
	}).updateControllerReference()
}

func newText(s string, size float64, ctr Controller) Text {
	return (&text{
		s:          s,
		size:       size,
		face:       text{}.newFace(size),
		Controller: ctr,
		debugImg:   newDebugImage(ctr),
	}).updateControllerReference()
}

/*
	Drawable
*/

func (t text) Draw(screen *ebiten.Image) {
	ebitentext.Draw(screen, t.s, t.face, &ebitentext.DrawOptions{
		DrawImageOptions: *t.Controller.DrawOption(),
		LayoutOptions:    ebitentext.LayoutOptions{LineSpacing: t.lineSpacing},
	})
}

func (t text) DebugDraw(screen *ebiten.Image, clr ...color.Color) {
	t.debugImg.Draw(screen, clr)
	t.Draw(screen)
}

/*
	embedController
*/

func (t *text) Align(a Align) Text {
	t.Controller.Align(a)
	return t
}

func (t *text) Move(x float64, y float64, replace ...bool) Text {
	t.Controller.Move(x, y, replace...)
	return t
}

func (t *text) Rotate(degree float64, replace ...bool) Text {
	t.Controller.Rotate(degree, replace...)
	return t
}

func (t *text) Scale(x float64, y float64, replace ...bool) Text {
	t.Controller.Scale(x, y, replace...)
	return t
}

func (t *text) updateControllerReference() Text {
	w, h := t.Bound()
	t.Controller.updateReference(w, h)
	t.debugImg.CleanCache()
	return t
}

/*
	Text
*/

func (t text) Copy(with ...Controller) Text {
	t.face = t.newFace(t.Size())

	if len(with) != 0 && with[0] != nil {
		t.Controller = with[0]
		return t.updateControllerReference()
	}

	return &t
}

func (t *text) SetColor(c color.Color) Text {
	t.color = c
	return t
}

func (t *text) SetLineSpacing(l float64) Text {
	t.lineSpacing = l
	return t.updateControllerReference()
}

func (t *text) SetText(text string) Text {
	t.s = text
	return t.updateControllerReference()
}

func (t *text) SetSize(size float64) Text {
	t.size = size
	t.face = t.newFace(size)
	return t.updateControllerReference()
}

func (t text) Bound() (w, h float64) {
	return ebitentext.Measure(t.s, t.face, t.lineSpacing)
}

func (t text) Color() color.Color {
	return t.color
}

func (f text) GetController() Controller {
	return f.Controller
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

func (t text) Vertexes() []Vector {
	return t.vertexes()
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
