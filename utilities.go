package ebitenpkg

import (
	sysimage "image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	_defaultDebugColor color.Color = color.RGBA{G: 100, A: 100}
)

func DebugImageFromImage(img sysimage.Image, clr ...color.Color) *ebiten.Image {
	bound := img.Bounds()
	return DebugImage(bound.Dx(), bound.Dy(), clr...)
}

func DebugImage(w, h int, clr ...color.Color) *ebiten.Image {
	debugColor := _defaultDebugColor
	if len(clr) != 0 && clr[0] != nil {
		debugColor = clr[0]
	}

	debugImg := ebiten.NewImage(w, h)
	debugImg.Fill(debugColor)
	return debugImg
}

type debugImage struct {
	ctr      Controller
	defaults *ebiten.Image
	cache    *ebiten.Image
	c        color.Color
}

func newDebugImage(ctr Controller) debugImage {
	return debugImage{
		ctr: ctr,
		c:   _defaultDebugColor,
	}
}

func (f *debugImage) Draw(screen *ebiten.Image, clr []color.Color) {
	if len(clr) != 0 {
		if f.cache == nil || !f.isColorEqual(f.c, clr[0]) {
			w, h := f.ctr.bound()
			f.cache = DebugImage(int(w), int(h), clr[0])
		}

		screen.DrawImage(f.cache, f.ctr.DrawOption())
		return
	}

	if f.defaults == nil {
		w, h := f.ctr.bound()
		f.defaults = DebugImage(int(w), int(h), _defaultDebugColor)
	}

	screen.DrawImage(f.defaults, f.ctr.DrawOption())
}

func (f *debugImage) CleanCache() {
	f.cache = nil
	f.defaults = nil
	f.c = _defaultDebugColor
}

func (debugImage) isColorEqual(a, b color.Color) bool {
	ar, ag, ab, aa := a.RGBA()
	br, bg, bb, ba := b.RGBA()
	return ar == br && ag == bg && ab == bb && aa == ba
}
