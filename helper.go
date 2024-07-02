package ebitenpkg

import (
	sysimage "image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	_defaultDebugColor color.Color = color.RGBA{G: 100, A: 100}
)

func DebugImageFromImage(img sysimage.Image, clr ...color.Color) (*ebiten.Image, color.Color) {
	bound := img.Bounds()
	return DebugImage(bound.Dx(), bound.Dy(), clr...)
}

func DebugImage(w, h int, clr ...color.Color) (*ebiten.Image, color.Color) {
	debugColor := _defaultDebugColor
	if len(clr) != 0 && clr[0] != nil {
		debugColor = clr[0]
	}

	debugImg := ebiten.NewImage(w, h)
	debugImg.Fill(debugColor)
	return debugImg, debugColor
}

type debugImage struct {
	ctr   Controller
	cache *ebiten.Image
	c     color.Color
}

func newDebugImage(ctr Controller) debugImage {
	return debugImage{
		ctr: ctr,
		c:   _defaultDebugColor,
	}
}

func (f *debugImage) Draw(screen *ebiten.Image, clr []color.Color) {
	if f.cache == nil || f.isColorCached(clr) {
		w, h := f.ctr.bound()
		f.cache, f.c = DebugImage(int(w), int(h), clr...)
	}

	screen.DrawImage(f.cache, f.ctr.DrawOption())
}

func (f *debugImage) CleanCache() {
	f.cache = nil
	f.c = _defaultDebugColor
}

func (f debugImage) isColorCached(clr []color.Color) bool {
	c := _defaultDebugColor
	if len(clr) != 0 && clr[0] != nil {
		c = clr[0]
	}

	return f.isColorEqual(f.c, c)
}

func (debugImage) isColorEqual(a, b color.Color) bool {
	ar, ag, ab, aa := a.RGBA()
	br, bg, bb, ba := b.RGBA()
	return ar == br && ag == bg && ab == bb && aa == ba
}
