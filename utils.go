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

func NewEbitenImageFromBounds(bounds func() (int, int), clr ...color.Color) *ebiten.Image {
	w, h := bounds()
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

func (i *canvas) DrawImageOn(top *ebiten.Image, x, y float64) *canvas {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(x, y)
	i.base.DrawImage(top, opt)

	return i
}

func (i *canvas) DrawRectOn(w, h int, clr color.Color, x, y float64) *canvas {
	i.DrawImageOn(NewEbitenImage(w, h, clr), x, y)

	return i
}

func (i *canvas) EbitenImage() *ebiten.Image {
	return i.base
}

func (i *canvas) Image() Image {
	return NewImage(i.base)
}
