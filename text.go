package ebitenpkg

import (
	"bytes"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yanun0323/pkg/logs"
)

type Text interface {
	Draw(screen *ebiten.Image)

	Align(align Align) Text
	Move(x, y float64, replace ...bool) Text
	Scale(x, y float64, replace ...bool) Text
	Rotate(angle float64, replace ...bool) Text
	Attach(parent Attachable) Text
	Detach() (parent Attachable)
	Debug(on ...bool) Text

	SetText(text string) Text
	SetSize(size float64) Text
	SetColor(color color.Color) Text
	SetLineSpacing(lineSpacing float64) Text
	SetFont(font []byte) Text

	Bounds() (width int, height int)
	Aligned() Align
	Moved() (x, y float64)
	Scaled() (x, y float64)
	Rotated() (angle float64)
	Debugged() bool

	Text() string
	Size() float64
	Color() color.Color
	LineSpacing() float64
	IsClick(x, y float64) bool
}

func NewText(text string, size float64) Text {
	return &eText{
		text:        newValue(text),
		size:        newValue(size),
		face:        newValue(newFace(size, DefaultFont())),
		color:       newValue(color.RGBA64{}),
		lineSpacing: newValue(0.0),
		font:        newValue(DefaultFont()),
	}
}

type eText struct {
	controller

	parent value[Attachable]
	debug  value[*ebiten.Image]

	text        value[string]
	size        value[float64]
	color       value[color.RGBA64]
	lineSpacing value[float64]
	face        value[text.Face]
	font        value[*text.GoTextFaceSource]
}

func (e *eText) Draw(screen *ebiten.Image) {
	opt := e.DrawOption()
	opt.ColorScale.ScaleWithColor(e.Color())

	text.Draw(screen, e.Text(), e.Face(), &text.DrawOptions{
		DrawImageOptions: *opt,
		LayoutOptions:    text.LayoutOptions{LineSpacing: e.LineSpacing()},
	})

	if debug := e.debug.Load(); debug != nil {
		screen.DrawImage(debug, opt)
	}
}

func (e *eText) Align(align Align) Text {
	e.controller.SetAlign(align)
	return e
}

func (e *eText) Move(x, y float64, replace ...bool) Text {
	e.controller.SetMove(x, y, replace...)
	return e
}

func (e *eText) Scale(x, y float64, replace ...bool) Text {
	e.controller.SetScale(x, y, replace...)
	return e
}

func (e *eText) Rotate(angle float64, replace ...bool) Text {
	e.controller.SetRotate(angle, replace...)
	return e
}

func (e *eText) Attach(parent Attachable) Text {
	e.parent.Store(parent)
	return e
}

func (e *eText) Detach() (parent Attachable) {
	e.parent.Store(nil)
	return e
}

func (e *eText) Debug(on ...bool) Text {
	debug := e.debug.Load()
	if len(on) != 0 && !on[0] {
		debug = nil
		return e
	}

	if debug != nil {
		return e
	}

	e.debug.Store(NewEbitenImageWith(e.Bounds, DefaultDebugColor()))
	return e
}

func (e *eText) SetText(text string) Text {
	e.text.Store(text)
	e.resetDebug()

	return e
}

func (e *eText) SetSize(size float64) Text {
	e.size.Store(size)
	e.face.Store(newFace(size, e.font.Load()))
	e.resetDebug()

	return e
}

func (e *eText) SetColor(clr color.Color) Text {
	r, g, b, a := clr.RGBA()
	e.color.Store(color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})

	return e
}

func (e *eText) SetLineSpacing(lineSpacing float64) Text {
	e.lineSpacing.Store(lineSpacing)
	e.resetDebug()

	return e
}

func (e *eText) SetFont(font []byte) Text {
	fs, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		logs.Fatalf("failed to load font: %v", err)
	}

	if fs == nil {
		logs.Fatalf("failed to load font")
	}

	e.font.Store(fs)
	e.face.Store(newFace(e.size.Load(), fs))
	e.resetDebug()

	return e
}

func (e *eText) Bounds() (width int, height int) {
	w, h := text.Measure(e.Text(), e.Face(), e.LineSpacing())
	return int(w), int(h)
}

func (e *eText) Aligned() Align {
	return e.controller.GetAlign()
}

func (e *eText) Moved() (x, y float64) {
	return e.controller.GetMove()
}

func (e *eText) Scaled() (x, y float64) {
	return e.controller.GetScale()
}

func (e *eText) Rotated() (angle float64) {
	return e.controller.GetRotate()
}

func (e *eText) Debugged() bool {
	return e.debug.Load() != nil
}

func (e *eText) Text() string {
	return e.text.Load()
}

func (e *eText) Face() text.Face {
	return e.face.Load()
}

func (e *eText) Size() float64 {
	return e.size.Load()
}

func (e *eText) Color() color.Color {
	return e.color.Load()
}

func (e *eText) LineSpacing() float64 {
	return e.lineSpacing.Load()
}

func (e *eText) IsClick(x, y float64) bool {
	w, h := e.Bounds()
	vertexes := getVertexes(float64(w), float64(h), e, e.parent.Load())
	return isInside(vertexes, Vector{X: x, Y: y})
}

func (e *eText) DrawOption() *ebiten.DrawImageOptions {
	w, h := e.Bounds()
	return getDrawOption(w, h, e.controller, 1, 1, e.parent.Load())
}

func (e *eText) resetDebug() {
	if debug := e.debug.Load(); debug != nil {
		e.debug.Store(nil)
	}
}

func newFace(size float64, fonts *text.GoTextFaceSource) text.Face {
	return &text.GoTextFace{
		Source: fonts,
		Size:   size,
	}
}
