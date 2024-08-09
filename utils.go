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

func NewEbitenRoundedImage(w, h int, round int, clr ...color.Color) *ebiten.Image {
	img := NewEbitenImage(w, h, clr...)
	for x := 0; x < round; x++ {
		// left-top
		ltX, ltY := round, round
		for y := 0; y < round; y++ {
			dx := x - ltX
			dy := y - ltY
			if dx*dx+dy*dy > round*round {
				img.Set(x, y, color.Transparent)
			}
		}

		// left-bottom
		lbX, lbY := round, h-round
		for y := h - round; y < h; y++ {
			dx := x - lbX
			dy := y - lbY
			if dx*dx+dy*dy > round*round {
				img.Set(x, y, color.Transparent)
			}
		}
	}

	for x := w - round; x < w; x++ {
		// right-top
		rtX, rtY := w-round, round
		for y := 0; y < round; y++ {
			dx := x - rtX
			dy := y - rtY
			if dx*dx+dy*dy > round*round {
				img.Set(x, y, color.Transparent)
			}
		}

		// right-bottom
		rbX, rbY := w-round, h-round
		for y := h - round; y < h; y++ {
			dx := x - rbX
			dy := y - rbY
			if dx*dx+dy*dy > round*round {
				img.Set(x, y, color.Transparent)
			}
		}
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

func (i *canvas) Image(children ...Attachable) Image {
	return NewImage(i.base, children...)
}
