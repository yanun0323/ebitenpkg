package ebito

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	_defaultDebugBorderWidth = 1
)

func DebugImageFromImage(img *ebiten.Image, borderWidth ...int) *ebiten.Image {
	bound := img.Bounds()
	return DebugImage(bound.Dx(), bound.Dy(), borderWidth...)
}

func DebugImage(w, h int, borderWidth ...int) *ebiten.Image {
	b := _defaultDebugBorderWidth
	if len(borderWidth) != 0 && borderWidth[0] >= 0 {
		b = borderWidth[0]
	}

	debugImg := ebiten.NewImage(w+b*2, h+b*2)
	bound := debugImg.Bounds()
	trailingX := bound.Dx() - b
	trailingY := bound.Dy() - b
	centerX := bound.Dx() / 2
	centerY := bound.Dy() / 2

	for x := 0; x < bound.Dx(); x++ {
		for y := 0; y < bound.Dy(); y++ {
			if x <= b || y <= b || x >= trailingX || y >= trailingY || x == centerX || y == centerY {
				debugImg.Set(x, y, color.White)
			}
		}
	}

	return debugImg
}
