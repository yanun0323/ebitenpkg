package ebitenpkg

import (
	sysimage "image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type image struct {
	Controller

	img *ebiten.Image
}

func NewImage(img sysimage.Image, a ...Align) Image {
	return (&image{
		img:        ebiten.NewImageFromImage(img),
		Controller: NewController(0, 0, a...),
	}).updateControllerReference()
}

func newImage(img sysimage.Image, ctr Controller) Image {
	return (&image{
		img:        ebiten.NewImageFromImage(img),
		Controller: ctr,
	}).updateControllerReference()
}

/*
	Drawable
*/

func (f image) Draw(screen *ebiten.Image) {
	screen.DrawImage(f.img, f.DrawOption())
}

/*
	DebugDrawable
*/

func (f image) DebugDraw(screen *ebiten.Image, clr ...color.Color) {
	f.Draw(screen)
	screen.DrawImage(DebugImageFromImage(f.img, clr...), f.DrawOption())
}

/*
	embedController
*/

func (f *image) Align(a Align) Image {
	f.Controller.Align(a)
	return f
}

func (f *image) Move(x float64, y float64, replace ...bool) Image {
	f.Controller.Move(x, y, replace...)
	return f
}

func (f *image) Rotate(degree float64, replace ...bool) Image {
	f.Controller.Rotate(degree, replace...)
	return f
}

func (f *image) Scale(x float64, y float64, replace ...bool) Image {
	f.Controller.Scale(x, y, replace...)
	return f
}

func (f *image) updateControllerReference() Image {
	f.Controller.updateReference(float64(f.img.Bounds().Dx()), float64(f.img.Bounds().Dy()))
	return f
}

/*
	Image
*/

func (f *image) Border(clr color.Color, width int) Image {
	if width <= 0 {
		return f
	}

	b := f.img.Bounds()
	zX, zY := width-1, width-1
	lX, lY := b.Dx()-width-1, b.Dy()-width-1
	for bx := 0; bx < b.Dx(); bx++ {
		for by := 0; by < b.Dy(); by++ {
			if bx <= zX || by <= zY || bx >= lX || by >= lY {
				f.img.Set(bx, by, clr)
			}
		}
	}

	return f
}

func (f image) Copy(with ...Controller) Image {
	b := f.img.Bounds()
	img := ebiten.NewImage(b.Dx(), b.Dy())
	img.DrawImage(f.img, nil)
	f.img = img

	if len(with) != 0 && with[0] != nil {
		f.Controller = with[0]
		return f.updateControllerReference()
	}

	return &f
}

func (f image) GetController() Controller {
	return f.Controller
}

func (f *image) ReplaceImage(img *ebiten.Image) Image {
	f.img = img
	return f.updateControllerReference()
}

func (f image) EbitenImage() *ebiten.Image {
	return f.img
}

func (f image) Vertexes() []vector {
	return f.vertexes()
}

/*
	private
*/
