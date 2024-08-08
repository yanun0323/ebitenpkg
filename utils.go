package ebitenpkg

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewEbitenImage(w, h int, clr ...color.Color) *ebiten.Image {
	img := ebiten.NewImage(w, h)
	if len(clr) != 0 {
		img.Fill(clr[0])
	}

	return img
}

func NewEbitenImageWith(b func() (int, int), clr ...color.Color) *ebiten.Image {
	w, h := b()
	return NewEbitenImage(w, h, clr...)
}

type canvas struct {
	base *ebiten.Image
}

func NewCanvas(w, h int, clr ...color.Color) *canvas {
	return &canvas{
		base: NewEbitenImage(w, h, clr...),
	}
}

func (i *canvas) DrawOn(top *ebiten.Image, x, y float64) *canvas {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(x, y)
	i.base.DrawImage(top, opt)

	return i
}

func (i *canvas) DrawImageOn(w, h int, clr color.Color, x, y float64) *canvas {
	i.DrawOn(NewEbitenImage(w, h, clr), x, y)

	return i
}

func (i *canvas) Image() *ebiten.Image {
	return i.base
}
