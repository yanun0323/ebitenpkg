package ebito

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	text        string
	font        *Font // cache
	*drawOption       // cache
}

func NewText(t string, size float64, a ...Align) *Text {
	return &Text{
		text:       t,
		font:       NewFont(size),
		drawOption: NewDrawOption(0, 0, a...),
	}
}

func (t Text) Text() string {
	return t.text
}

func (t Text) Size() (w, h float64) {
	return text.Measure(t.text, t.font.GoXFace(), t.font.LinsSpace())
}

func (t Text) Copy() *Text {
	opt := *t.drawOption
	t.drawOption = &opt
	t.font = t.font.Copy()
	return &t
}

func (t *Text) SetText(text string) {
	t.text = text
}

func (t *Text) WithText(text string) *Text {
	tt := t.Copy()
	tt.text = text
	return tt
}

func (t Text) Draw(screen *ebiten.Image) {
	w, h := t.Size()
	t.drawOption.updateReference(w, h)
	text.Draw(screen, t.text, t.font.GoXFace(), &text.DrawOptions{
		DrawImageOptions: *t.drawOption.Option(),
		LayoutOptions:    t.font.LayoutOption(true),
	})
}

func (t Text) DebugDraw(screen *ebiten.Image, borderWidth ...int) {
	t.Draw(screen)
	w, h := t.Size()
	screen.DrawImage(DebugImage(int(w), int(h), borderWidth...), t.debugOption(borderWidth...))
}

// Extension of Font

func (t Text) Font() *Font {
	return t.font
}

func (t *Text) FontGoXFace() *text.GoXFace {
	return t.font.face
}

func (t *Text) FontSize() float64 {
	return t.font.size
}

func (t *Text) FontColor() color.Color {
	return t.font.color
}

func (t *Text) FontLinsSpace() float64 {
	return t.font.lineSpacing
}

func (t *Text) WithFontColor(c color.Color) *Text {
	tt := t.Copy()
	tt.font.WithColor(c)
	return tt
}

func (t *Text) WithFontLinsSpace(spacing float64) *Text {
	tt := t.Copy()
	tt.font.WithLinsSpace(spacing)
	return tt
}

// Extension of DrawOption

func (t Text) WithMovement(x, y float64) *Text {
	tt := t.Copy()
	tt.drawOption = tt.drawOption.withMovement(x, y)
	return tt
}

func (t Text) WithScaleRatio(x, y float64) *Text {
	tt := t.Copy()
	tt.drawOption = tt.drawOption.withScaleRatio(x, y)
	return tt
}

func (t Text) WithRotation(degree float64) *Text {
	tt := t.Copy()
	tt.drawOption = tt.drawOption.withRotation(degree)
	return tt
}

func (t Text) WithAlignment(a Align) *Text {
	tt := t.Copy()
	tt.drawOption = tt.drawOption.withAlignment(a)
	return tt
}
