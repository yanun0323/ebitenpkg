package ebitenpkg

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	gofont "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	_dpi = 72
)

var (
	_defaultFont             = Font{face: text.NewGoXFace(nil)}
	_defaultFontSize float64 = 24
)

type Font struct {
	face           *text.GoXFace
	size           float64
	color          color.Color
	lineSpacing    float64
	primaryAlign   Align
	secondaryAlign Align
}

func NewFont(size ...float64) *Font {
	s := _defaultFontSize
	if len(size) != 0 && size[0] >= 0 {
		s = size[0]
	}

	f, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return _defaultFont.Copy()
	}

	ft, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    s,
		DPI:     _dpi,
		Hinting: gofont.HintingFull,
	})
	if err != nil {
		return _defaultFont.Copy()
	}

	return &Font{
		face: text.NewGoXFace(ft),
	}
}

func (f Font) GoXFace() *text.GoXFace {
	return f.face
}

func (f Font) Size() float64 {
	return f.size
}

func (f Font) Color() color.Color {
	return f.color
}

func (f Font) LinsSpace() float64 {
	return f.lineSpacing
}

func (f Font) PrimaryAlign() Align {
	return f.primaryAlign
}

func (f Font) SecondaryAlign() Align {
	return f.secondaryAlign
}

func (f Font) Copy() *Font {
	return &f
}

func (f *Font) SetSize(size float64) {
	f.size = size
	f.face = NewFont(size).face
}

func (f *Font) WithColor(c color.Color) *Font {
	ff := f.Copy()
	ff.color = c
	return ff
}

func (f *Font) WithLinsSpace(spacing float64) *Font {
	ff := f.Copy()
	ff.lineSpacing = spacing
	return ff
}

func (f *Font) WithPrimaryAlign(a Align) *Font {
	ff := f.Copy()
	ff.primaryAlign = a
	return ff
}

func (f *Font) WithSecondaryAlign(a Align) *Font {
	ff := f.Copy()
	ff.secondaryAlign = a
	return ff
}

func (f Font) LayoutOption(ignoreAlign ...bool) text.LayoutOptions {
	if len(ignoreAlign) != 0 && ignoreAlign[0] {
		return text.LayoutOptions{LineSpacing: f.lineSpacing}
	}
	return text.LayoutOptions{
		LineSpacing:    f.lineSpacing,
		PrimaryAlign:   f.primaryAlign.TextAlign(),
		SecondaryAlign: f.secondaryAlign.TextAlign(),
	}
}
