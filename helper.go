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
