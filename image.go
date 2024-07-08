package ebitenpkg

import (
	sysimage "image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type image struct {
	ctr             controller
	debugImageCache debugCache
	img             *ebiten.Image
}

func NewImage(img sysimage.Image, a ...Align) Image {
	return &image{
		ctr: newController(a...),
		img: ebiten.NewImageFromImage(img),
	}
}

/*
	Drawable
*/

func (im image) Draw(screen *ebiten.Image) {
	screen.DrawImage(im.img, im.DrawOption())
}

func (im image) DebugDraw(screen *ebiten.Image, clr ...color.Color) {
	opt := im.DrawOption()
	debugImg := im.debugImageCache.Image(im.img.Bounds().Dx(), im.img.Bounds().Dy(), clr...)
	screen.DrawImage(im.img, opt)
	screen.DrawImage(debugImg, opt)
}

/*
	Controllable
*/

func (im *image) Align(a Align) Image {
	im.ctr.Align(a)
	return im
}

func (im *image) Move(x, y float64, replace ...bool) Image {
	im.ctr.Move(x, y, replace...)
	return im
}

func (im *image) Rotate(degree float64, replace ...bool) Image {
	im.ctr.Rotate(degree, replace...)
	return im
}

func (im *image) Scale(x, y float64, replace ...bool) Image {
	im.ctr.Scale(x, y, replace...)
	return im
}

func (im image) Aligned() Align {
	return im.ctr.Aligned()
}

func (im image) Moved() (x, y float64) {
	return im.ctr.Moved()
}

func (im image) Rotated() float64 {
	return im.ctr.Rotated()
}

func (im image) Scaled() (x, y float64) {
	return im.ctr.Scaled()
}

func (im image) DrawOption() *ebiten.DrawImageOptions {
	w, h := im.Bounds()
	return getDrawOption(w, h, im.ctr)
}

func (im image) Bounds() (w, h float64) {
	b := im.img.Bounds()
	return float64(b.Dx()), float64(b.Dy())
}

func (im image) Barycenter() (x, y float64) {
	return im.ctr.Moved()
}

/*
	Image
*/

func (im *image) Border(clr color.Color, width int) Image {
	if width <= 0 {
		return im
	}

	b := im.img.Bounds()
	zX, zY := width-1, width-1
	lX, lY := b.Dx()-width-1, b.Dy()-width-1
	for bx := 0; bx < b.Dx(); bx++ {
		for by := 0; by < b.Dy(); by++ {
			if bx <= zX || by <= zY || bx >= lX || by >= lY {
				im.img.Set(bx, by, clr)
			}
		}
	}

	return im
}

func (im image) Copy() Image {
	cp := im.copy()
	return &cp
}

func (im *image) ReplaceImage(img *ebiten.Image) Image {
	im.img = img
	im.debugImageCache.Clean()
	return im
}

func (im image) EbitenImage() *ebiten.Image {
	return im.img
}

/*
	Private
*/

func (im image) copy() image {
	b := im.img.Bounds()
	img := ebiten.NewImage(b.Dx(), b.Dy())
	img.DrawImage(im.img, nil)
	im.img = img
	im.debugImageCache = im.debugImageCache.Copy()

	return im
}
